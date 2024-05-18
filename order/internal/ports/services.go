package ports

import (
	"context"
	"orderService/internal/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *models.Order) (string, error)
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	ChangeOrderStatus(ctx context.Context, id string, status string) (string, error)
	CancelOrder(ctx context.Context, id string) (string, error)
}

type KitchenServiceClientProxy interface {
	ProcessOrder(ctx context.Context, order *models.Order) error
	ChangeOrderStatus(ctx context.Context, orderId, status string) error
}
