package ports

import (
	"context"
	"kitchenService/internal/models"
)

type OrderProxy interface {
	ChangeOrderStatus(ctx context.Context, orderId, status string) error
}

type KitchenService interface {
	ProcessOrder(ctx context.Context, order *models.Order) (string, *models.OrderStatus, error)
	ChangeOrderStatus(ctx context.Context, orderId string, status *models.OrderStatus) error
}
