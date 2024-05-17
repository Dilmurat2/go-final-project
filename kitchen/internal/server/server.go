package server

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"kitchenService/internal/models"
	"kitchenService/internal/ports"
	"kitchenService/pkg/helpers"
	"kitchenService/proto/kitchen"
	"kitchenService/proto/order"
)

type Server struct {
	kitchen.UnimplementedKitchenServiceServer
	ports.KitchenService
}

func NewServer(kitchenService ports.KitchenService) *Server {
	return &Server{KitchenService: kitchenService}
}

func (s Server) ProcessOrder(ctx context.Context, orderPb *order.Order) (*kitchen.ProcessOrderResponse, error) {
	order := helpers.OrderPbToModel(orderPb)
	orderId, status, err := s.KitchenService.ProcessOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &kitchen.ProcessOrderResponse{
		Id:     orderId,
		Status: string(*status),
	}, nil
}

func (s Server) ChangeOrderStatus(ctx context.Context, request *order.ChangeOrderStatusRequest) (*emptypb.Empty, error) {

	fmt.Println(request.GetId())
	orderId := request.GetId()
	orderStatus := models.OrderStatus(request.GetStatus())
	err := s.KitchenService.ChangeOrderStatus(ctx, orderId, &orderStatus)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
