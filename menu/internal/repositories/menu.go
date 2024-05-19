package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"menu/internal/config"
	"menu/internal/models"
	"menu/internal/ports"
	"time"
)

const (
	MenuCollection = "Menu"
	MenuDatabase   = "Menu"
)

type menuRepository struct {
	mongo *mongo.Client
}

func NewMenuRepository(cfg *config.Config) (ports.MenuRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%v:%v@%v:%v",
			cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port,
		)).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)),
	)
	if err != nil {
		return nil, err
	}

	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, err
	}

	return &menuRepository{mongo: client}, nil
}

func (m menuRepository) AddItem(ctx context.Context, menuID string, item *models.Item) (*models.Menu, error) {
	var menu models.Menu
	filter := bson.M{"_id": menuID}
	update := bson.M{"$push": bson.M{"items": item}}
	_, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error adding item to menu: %w", err)
	}
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).FindOne(ctx, filter).Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("error getting menu: %w", err)
	}
	return &menu, nil
}

func (m menuRepository) DeleteItem(ctx context.Context, menuID string, itemID string) (*models.Menu, error) {
	var menu models.Menu

	filter := bson.M{"_id": menuID}
	update := bson.M{"$pull": bson.M{"items": bson.M{"_id": itemID}}}
	_, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error deleting item from menu: %w", err)
	}
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).FindOne(ctx, filter).Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("error getting menu: %w", err)
	}
	return &menu, nil
}

func (m menuRepository) GetAll(ctx context.Context) ([]*models.Menu, error) {
	var menus []*models.Menu
	cursor, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting menus: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var menu models.Menu
		if err := cursor.Decode(&menu); err != nil {
			return nil, fmt.Errorf("error decoding menu: %w", err)
		}
		menus = append(menus, &menu)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating cursor: %w", err)
	}

	return menus, nil
}

func (m menuRepository) GetByID(ctx context.Context, id string) (*models.Menu, error) {
	var menu models.Menu
	filter := bson.M{"_id": id}

	err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).FindOne(ctx, filter).Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("error getting menu: %w", err)
	}
	return &menu, nil
}
func (m menuRepository) Update(ctx context.Context, menu *models.Menu) (*models.Menu, error) {
	filter := bson.M{"_id": menu.ID}

	update := bson.M{
		"$set": bson.M{
			"name":        menu.Name,
			"description": menu.Description,
			"updated_at":  time.Now().Format(time.RFC3339),
			"is_active":   menu.IsActive,
		},
	}

	_, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error updating menu: %w", err)
	}

	return menu, nil
}

func (m menuRepository) Delete(ctx context.Context, id string) error {
	_, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return fmt.Errorf("error deleting menu: %w", err)
	}
	return nil
}

func (m menuRepository) Create(ctx context.Context, menu *models.Menu) (*models.Menu, error) {
	result, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).InsertOne(ctx, menu)
	if err != nil {
		return nil, fmt.Errorf("error inserting menu: %w", err)
	}

	menu.ID = result.InsertedID.(string)
	return menu, nil
}
