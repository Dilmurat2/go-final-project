package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"orderService/config"
	"orderService/internal/models"
)

type OrderRepository struct {
	mongo *mongo.Client
}

func NewOrderRepository(cfg *config.Config) (*OrderRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%v:%v@%v:%v",
			cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port,
		)).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %v", err)
	}

	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %v", err)
	}

	return &OrderRepository{mongo: client}, nil
}

func (o *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) (string, error) {
	id, err := o.mongo.Database("orders").Collection("orders").InsertOne(ctx, order)
	if err != nil {
		return "", fmt.Errorf("failed to insert order: %v", err)
	}

	return id.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (o *OrderRepository) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	order := new(models.Order)
	if err = o.mongo.Database("orders").Collection("orders").FindOne(ctx, bson.M{"_id": objectID}).Decode(order); err != nil {
		return nil, fmt.Errorf("failed to get order: %v", err)
	}

	return order, nil
}

func (o *OrderRepository) ChangeOrderStatus(ctx context.Context, id, status string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("invalid id: %v", err)
	}

	update := bson.M{"$set": bson.M{"status": status}}
	if _, err = o.mongo.Database("orders").Collection("orders").UpdateOne(ctx, bson.M{"_id": objectID}, update); err != nil {
		return "", fmt.Errorf("failed to cancel order: %v", err)
	}

	return id, nil
}
