package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "in_progress"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusCompleted OrderStatus = "completed"
)

type Order struct {
	ID        string      `json:"id" bson:"_id,omitempty"`
	Items     []*Item     `json:"items" bson:"items"`
	Status    OrderStatus `json:"status" bson:"status"`
	CreatedAt string      `json:"created_at" bson:"created_at"`
	UpdatedAt string      `json:"updated_at" bson:"updated_at"`
	DeletedAt string      `json:"deleted_at" bson:"deleted_at"`
}

func NewOrder(items []*Item) *Order {
	order := Order{
		ID:        primitive.NewObjectID().Hex(),
		Status:    OrderStatusPending,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	for _, item := range items {
		order.Items = append(order.Items, item)
	}
	return &order
}
