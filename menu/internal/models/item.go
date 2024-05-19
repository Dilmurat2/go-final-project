package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Item struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float32 `json:"price" bson:"price"`
	Weight      float32 `json:"weight" bson:"weight"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
	UpdatedAt   string  `json:"updated_at" bson:"updated_at"`
	DeletedAt   string  `json:"deleted_at" bson:"deleted_at"`
	IsActive    bool    `json:"is_active" bson:"is_active"`
}

func NewItem(name, desc string, price float32, weight float32) *Item {
	return &Item{
		ID:          primitive.NewObjectID().Hex(),
		Description: desc,
		Name:        name,
		Price:       price,
		Weight:      weight,
		CreatedAt:   time.Now().Format(time.RFC3339),
		IsActive:    true,
	}
}
