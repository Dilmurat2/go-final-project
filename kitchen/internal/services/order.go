package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"kitchenService/proto/order"
	"log"
)

type orderProxy struct {
	client order.OrderServiceClient
}

func NewOrderProxy() *orderProxy {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := order.NewOrderServiceClient(conn)

	return &orderProxy{client: client}
}

func (op *orderProxy) ChangeOrderStatus(ctx context.Context, orderId, status string) error {
	_, err := op.client.ChangeOrderStatus(ctx, &order.ChangeOrderStatusRequest{
		Id:     orderId,
		Status: status,
	})
	if err != nil {
		return fmt.Errorf("could not change order status: %v", err)
	}
	return nil
}
