package services_test

import (
	"context"
	"errors"
	"kitchenService/internal/models"
	"kitchenService/internal/services"
	mock_ports "kitchenService/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_ports.NewMockKitchenRepository(ctrl)
	mockOrderClient := mock_ports.NewMockOrderProxy(ctrl)

	kitchenService := services.NewKitchenService(mockRepo, mockOrderClient)

	order := &models.Order{
		ID:     "123",
		Items:  []models.Item{{ID: 1, Name: "Pizza", Price: 12.99}},
		Status: models.OrderStatusPending,
	}
	orderStatus := models.OrderStatusPending

	mockRepo.EXPECT().ProcessOrder(gomock.Any(), order).Return("order123", &orderStatus, nil)

	orderID, status, err := kitchenService.ProcessOrder(context.Background(), order)

	assert.NoError(t, err)
	assert.Equal(t, "order123", orderID)
	assert.Equal(t, orderStatus, *status)
}

func TestChangeOrderStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_ports.NewMockKitchenRepository(ctrl)
	mockOrderClient := mock_ports.NewMockOrderProxy(ctrl)

	kitchenService := services.NewKitchenService(mockRepo, mockOrderClient)

	orderID := "123"
	status := models.OrderStatusCompleted

	mockOrderClient.EXPECT().ChangeOrderStatus(gomock.Any(), orderID, string(status)).Return(nil)
	mockRepo.EXPECT().ChangeOrderStatus(gomock.Any(), orderID, &status).Return(nil)

	err := kitchenService.ChangeOrderStatus(context.Background(), orderID, &status)

	assert.NoError(t, err)
}

func TestChangeOrderStatus_OrderClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_ports.NewMockKitchenRepository(ctrl)
	mockOrderClient := mock_ports.NewMockOrderProxy(ctrl)

	kitchenService := services.NewKitchenService(mockRepo, mockOrderClient)

	orderID := "123"
	status := models.OrderStatusCompleted

	mockOrderClient.EXPECT().ChangeOrderStatus(gomock.Any(), orderID, string(status)).Return(errors.New("order service error"))

	err := kitchenService.ChangeOrderStatus(context.Background(), orderID, &status)

	assert.Error(t, err)
	assert.Equal(t, "order service error", err.Error())
}
