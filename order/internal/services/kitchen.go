package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"orderService/config"
	"orderService/internal/models"
	"orderService/pkg/helpers"
	kitchen "orderService/proto/kitchen"
	order_v1 "orderService/proto/order"
)

type kitchenProxy struct {
	client kitchen.KitchenServiceClient
}

func NewKitchenProxy(cfg *config.Config) *kitchenProxy {
	conn, err := grpc.Dial(cfg.KitchenServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := kitchen.NewKitchenServiceClient(conn)

	return &kitchenProxy{client: client}
}

func (kc *kitchenProxy) ProcessOrder(ctx context.Context, order *models.Order) error {
	orderReq := helpers.OrderModelToProtobuf(order)
	_, err := kc.client.ProcessOrder(ctx, orderReq)
	if err != nil {
		return fmt.Errorf("could not process order: %v", err)
	}
	return nil
}

func (kc *kitchenProxy) ChangeOrderStatus(ctx context.Context, orderId, status string) error {
	_, err := kc.client.ChangeOrderStatus(ctx, &order_v1.ChangeOrderStatusRequest{
		Id:     orderId,
		Status: status,
	})
	fmt.Println(orderId)
	if err != nil {
		return fmt.Errorf("could not change order status: %v", err)
	}
	return nil

}
