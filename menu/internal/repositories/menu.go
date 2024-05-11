package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"menu/internal/config"
	"menu/internal/models"
	"menu/internal/ports"
)

const (
	MenuCollection = "Menu"
	MenuDatabase   = "Menu"
)

type menuRepository struct {
	mongo *mongo.Client
}

func (m menuRepository) GetAll(ctx context.Context) ([]*models.Menu, error) {
	var menus []*models.Menu
	cursor, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting menus: %w", err)
	}
	for cursor.Next(ctx) {
		var menu models.Menu
		if err := cursor.Decode(&menu); err != nil {
			return nil, fmt.Errorf("error decoding menu: %w", err)
		}
		menus = append(menus, &menu)
	}
	return menus, nil
}

func (m menuRepository) GetByID(ctx context.Context, id string) (*models.Menu, error) {
	var menu models.Menu
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error converting id to object id: %w", err)
	}
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		FindOne(ctx, bson.D{{Key: "_id", Value: oid}}).Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("error getting menu: %w", err)
	}
	return &menu, nil
}

func (m menuRepository) Update(ctx context.Context, menu *models.Menu) (*models.Menu, error) {
	var updatedMenu models.Menu
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		FindOneAndUpdate(ctx, bson.D{{Key: "_id", Value: menu.ID}}, bson.D{{Key: "$set", Value: menu}}, opts).Decode(&updatedMenu)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("menu not found: %w", err)
		}
		return nil, fmt.Errorf("error updating menu: %w", err)
	}
	return &updatedMenu, nil
}

func (m menuRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting id to object id: %w", err)
	}
	_, err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		DeleteOne(ctx, bson.D{{Key: "_id", Value: oid}})
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
