package server

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"kitchenService/internal/models"
	"kitchenService/internal/ports"
	"kitchenService/pkg/helpers"
	kitchenv1 "kitchenService/proto/v1"
)

type Server struct {
	kitchenv1.UnimplementedKitchenServiceServer
	ports.KitchenService
}

func NewServer(kitchenService ports.KitchenService) *Server {
	return &Server{KitchenService: kitchenService}
}

func (s Server) ProcessOrder(ctx context.Context, orderPb *kitchenv1.Order) (*kitchenv1.ProcessOrderResponse, error) {
	order := helpers.OrderPbToModel(orderPb)
	orderId, status, err := s.KitchenService.ProcessOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &kitchenv1.ProcessOrderResponse{
		Id:     orderId,
		Status: string(*status),
	}, nil
}

func (s Server) ChangeOrderStatus(ctx context.Context, request *kitchenv1.ChangeOrderStatusRequest) (*emptypb.Empty, error) {
	orderId := request.GetOrderId()
	orderStatus := models.OrderStatus(request.GetOrderStatus())
	err := s.KitchenService.ChangeOrderStatus(ctx, orderId, &orderStatus)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
