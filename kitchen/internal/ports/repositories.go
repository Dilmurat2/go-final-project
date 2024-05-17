package ports

import (
	"context"
	"kitchenService/internal/models"
)

type KitchenRepository interface {
	ProcessOrder(ctx context.Context, order *models.Order) (string, *models.OrderStatus, error)
	ChangeOrderStatus(ctx context.Context, orderId string, status *models.OrderStatus) error
}
