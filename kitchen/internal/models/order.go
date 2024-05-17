package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "in_progress"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusCompleted OrderStatus = "completed"
)

type Order struct {
	ID     string      `json:"id"`
	Items  []Item      `json:"items"`
	Status OrderStatus `json:"status"`
}
