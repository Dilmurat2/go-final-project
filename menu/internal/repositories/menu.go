package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m menuRepository) GetItems(ctx context.Context, menuID string) (*[]models.Item, error) {
	panic("Implement me")
}

func (m menuRepository) AddItem(ctx context.Context, menuID string, item *models.Item) (*models.Menu, error) {
	var menu models.Menu
	oid, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return nil, fmt.Errorf("error converting id to object id: %w", err)
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$push": bson.M{"items": item}}
	_, err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error adding item to menu: %w", err)
	}
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		FindOne(ctx, bson.D{{Key: "_id", Value: menuID}}).Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("error getting menu: %w", err)
	}
	return &menu, nil
}

func (m menuRepository) DeleteItem(ctx context.Context, menuID string, itemID string) (*models.Menu, error) {
	menuObjID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return nil, err
	}

	itemObjID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": menuObjID}
	update := bson.M{"$pull": bson.M{"items": bson.M{"_id": itemObjID}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedMenu models.Menu
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedMenu)
	if err != nil {
		return nil, err
	}

	return &updatedMenu, nil
}

func (m menuRepository) GetItem(ctx context.Context, menuID string, itemID string) (*models.Item, error) {
	menuObjID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return nil, err
	}
	itemObjID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": menuObjID, "items._id": itemObjID}
	projection := bson.M{"items.$": 1}

	var menu models.Menu
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&menu)
	if err != nil {
		return nil, err
	}

	if len(menu.Items) > 0 {
		return &menu.Items[0], nil
	}

	return nil, mongo.ErrNoDocuments
}

func (m menuRepository) UpdateItem(ctx context.Context, menuID string, itemID string, updatedItem *models.Item) (*models.Item, error) {
	menuObjID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return nil, err
	}

	itemObjID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}

	updatedItem.UpdatedAt = time.Now().Format(time.RFC3339)
	filter := bson.M{"_id": menuObjID, "items._id": itemObjID}
	update := bson.M{
		"$set": bson.M{
			"items.$.name":       updatedItem.Name,
			"items.$.price":      updatedItem.Price,
			"items.$.weight":     updatedItem.Weight,
			"items.$.updated_at": updatedItem.UpdatedAt,
			"items.$.is_active":  updatedItem.IsActive,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedMenu models.Menu
	err = m.mongo.Database(MenuDatabase).Collection(MenuCollection).FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedMenu)
	if err != nil {
		return nil, err
	}

	for _, item := range updatedMenu.Items {
		if item.ID == itemID {
			return &item, nil
		}
	}

	return nil, mongo.ErrNoDocuments
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

	err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).
		FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("error getting menu: %w", err)
	}
	return &menu, nil
}

func (m menuRepository) Update(ctx context.Context, menu *models.Menu) (*models.Menu, error) {
	filter := bson.M{"_id": menu.ID}
	update := bson.M{"$set": menu}
	_, err := m.mongo.Database(MenuDatabase).Collection(MenuCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error updating menu: %w", err)
	}
	return menu, nil
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
