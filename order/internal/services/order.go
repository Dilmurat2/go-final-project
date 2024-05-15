package services

import (
	"context"
	"orderService/internal/models"
	"orderService/internal/repositories"
)

type OrderService struct {
	orderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (o *OrderService) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	newOrder := models.NewOrder(order.Items, order.Status)
	return o.orderRepository.CreateOrder(ctx, newOrder)
}

func (o *OrderService) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return o.orderRepository.GetOrder(ctx, id)
}

func (o *OrderService) CancelOrder(ctx context.Context, id, status string) (string, error) {
	return o.orderRepository.ChangeOrderStatus(ctx, id, status)
}
