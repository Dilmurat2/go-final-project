package services

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"orderService/internal/models"
	"orderService/internal/services"
	"orderService/tests/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderService_CreateOrder_KitchenServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepository := mock.NewMockOrderRepository(ctrl)
	mockKitchenServiceClientProxy := mock.NewMockKitchenServiceClientProxy(ctrl)

	orderService := services.NewOrderService(mockOrderRepository, mockKitchenServiceClientProxy)

	order := &models.Order{Items: []*models.Item{
		{ID: "1", Name: "item1", Price: 10},
		{ID: "2", Name: "item2", Price: 20},
	},
	}

	mockKitchenServiceClientProxy.EXPECT().ProcessOrder(gomock.Any(), gomock.Any()).Return(errors.New("kitchen service error"))

	_, err := orderService.CreateOrder(context.Background(), order)

	assert.Error(t, err)
	assert.Equal(t, "kitchen service error", err.Error())
}

func TestOrderService_CancelOrder_OrderNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepository := mock.NewMockOrderRepository(ctrl)
	mockKitchenServiceClientProxy := mock.NewMockKitchenServiceClientProxy(ctrl)

	orderService := services.NewOrderService(mockOrderRepository, mockKitchenServiceClientProxy)

	mockOrderRepository.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Return(nil, errors.New("order not found"))

	_, err := orderService.CancelOrder(context.Background(), "orderID")

	assert.Error(t, err)
	assert.Equal(t, "order not found", err.Error())
}

func TestOrderService_CancelOrder_OrderCannotBeCancelled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepository := mock.NewMockOrderRepository(ctrl)
	mockKitchenServiceClientProxy := mock.NewMockKitchenServiceClientProxy(ctrl)

	orderService := services.NewOrderService(mockOrderRepository, mockKitchenServiceClientProxy)

	order := &models.Order{Status: models.OrderStatusCompleted, CreatedAt: "2021-01-01T00:00:00Z"}

	mockOrderRepository.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Return(order, nil)

	_, err := orderService.CancelOrder(context.Background(), "orderID")

	assert.Error(t, err)
	assert.Equal(t, "order cannot be cancelled", err.Error())
}
