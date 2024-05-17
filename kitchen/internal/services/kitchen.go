package services

import (
	"context"
	"kitchenService/internal/models"
	"kitchenService/internal/ports"
)

type KitchenService struct {
	kitchenRepository ports.KitchenRepository
	orderClient       *orderProxy
}

func NewKitchenService(kitchenRepository ports.KitchenRepository, oc *orderProxy) ports.KitchenService {
	return &KitchenService{kitchenRepository: kitchenRepository,
		orderClient: oc}
}

func (k KitchenService) ProcessOrder(ctx context.Context, order *models.Order) (string, *models.OrderStatus, error) {
	return k.kitchenRepository.ProcessOrder(ctx, order)
}

func (k KitchenService) ChangeOrderStatus(ctx context.Context, orderId string, status *models.OrderStatus) error {
	err := k.orderClient.ChangeOrderStatus(ctx, orderId, string(*status))
	if err != nil {
		return err
	}
	return k.kitchenRepository.ChangeOrderStatus(ctx, orderId, status)
}
