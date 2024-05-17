package services

import (
	"context"
	"kitchenService/internal/models"
	"kitchenService/internal/ports"
)

type KitchenService struct {
	kitchenRepository ports.KitchenRepository
}

func NewKitchenService(kitchenRepository ports.KitchenRepository) ports.KitchenService {
	return &KitchenService{kitchenRepository: kitchenRepository}
}

func (k KitchenService) ProcessOrder(ctx context.Context, order *models.Order) (string, *models.OrderStatus, error) {
	return k.kitchenRepository.ProcessOrder(ctx, order)
}

func (k KitchenService) ChangeOrderStatus(ctx context.Context, orderId string, status *models.OrderStatus) error {
	return k.ChangeOrderStatus(ctx, orderId, status)
}
