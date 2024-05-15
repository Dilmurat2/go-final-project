package models

import "time"

type Order struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Items     []*Item `json:"items" bson:"items"`
	Status    string  `json:"status" bson:"status"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	DeletedAt string  `json:"deleted_at" bson:"deleted_at"`
}

func NewOrder(items []*Item, status string) *Order {
	var order Order
	order.Status = status
	for _, item := range items {
		order.Items = append(order.Items, item)
	}
	order.CreatedAt = time.Now().Format(time.RFC3339)
	return &order
}
