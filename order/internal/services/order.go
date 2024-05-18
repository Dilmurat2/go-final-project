package services

import (
	"context"
	"fmt"
	"orderService/internal/models"
	"orderService/internal/ports"
	"orderService/pkg/app_errors"
	"orderService/pkg/helpers"
)

type OrderService struct {
	orderRepository ports.OrderRepository
	kitchenClient   *kitchenProxy
}

func NewOrderService(or ports.OrderRepository, kc *kitchenProxy) ports.OrderService {
	return &OrderService{
		orderRepository: or,
		kitchenClient:   kc,
	}
}
func (o *OrderService) CreateOrder(ctx context.Context, order *models.Order) (string, error) {
	newOrder := models.NewOrder(order.Items)
	if err := o.kitchenClient.ProcessOrder(ctx, newOrder); err != nil {
		fmt.Printf("failed to process order: %v", err)
	}
	return o.orderRepository.CreateOrder(ctx, newOrder)
}

func (o *OrderService) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return o.orderRepository.GetOrder(ctx, id)
}

func (o *OrderService) ChangeOrderStatus(ctx context.Context, id, status string) (string, error) {
	return o.orderRepository.ChangeOrderStatus(ctx, id, status)
}

func (o *OrderService) CancelOrder(ctx context.Context, id string) (string, error) {
	order, err := o.orderRepository.GetOrder(ctx, id)
	if err != nil {
		return "", err
	}
	t, err := helpers.CalculateTimeSinceCreation(order.CreatedAt)
	if err != nil {
		return "", fmt.Errorf("failed to calculate time since creation: %v", err)
	}
	if order.Status != models.OrderStatusPending && t >= 5 {
		return "", app_errors.ErrCantCancelOrder
	}
	if err := o.kitchenClient.ChangeOrderStatus(ctx, id, string(models.OrderStatusCancelled)); err != nil {
		return "", err
	}
	return o.orderRepository.ChangeOrderStatus(ctx, id, string(models.OrderStatusCancelled))
}