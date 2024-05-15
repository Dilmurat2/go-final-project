package server

import (
	"context"
	"orderService/internal/services"
	"orderService/pkg/helpers"
	order_v1 "orderService/proto/v1"
)

type Server struct {
	order_v1.UnimplementedOrderServiceServer
	orderService *services.OrderService
}

func NewServer(orderService *services.OrderService) *Server {
	return &Server{
		orderService: orderService,
	}
}

func (s *Server) CreateOrder(ctx context.Context, req *order_v1.Order) (*order_v1.CreateOrderResponse, error) {
	order := helpers.OrderProtobufToModel(req)
	id, err := s.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &order_v1.CreateOrderResponse{
		Id:     id,
		Status: "Created",
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *order_v1.GetOrderRequest) (*order_v1.Order, error) {
	order, err := s.orderService.GetOrder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return helpers.OrderModelToProtobuf(order), nil
}

func (s *Server) ChangeOrderStatus(ctx context.Context, req *order_v1.ChangeOrderStatusRequest) (*order_v1.ChangeOrderStatusResponse, error) {
	id, err := s.orderService.ChangeOrderStatus(ctx, req.GetId(), req.GetStatus())
	if err != nil {
		return nil, err
	}
	return &order_v1.ChangeOrderStatusResponse{
		Id:     id,
		Status: "Updated",
	}, nil
}
