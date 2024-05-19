package server

import (
	"context"
	"errors"
	"fmt"
	"orderService/internal/ports"
	"orderService/pkg/app_errors"
	"orderService/pkg/helpers"
	order_v1 "orderService/proto/order"
)

type Server struct {
	order_v1.UnimplementedOrderServiceServer
	orderService ports.OrderService
}

func NewServer(orderService ports.OrderService) *Server {
	return &Server{
		orderService: orderService,
	}
}

func (s *Server) CreateOrder(ctx context.Context, req *order_v1.Order) (*order_v1.CreateOrderResponse, error) {
	order := helpers.OrderProtobufToModel(req)
	fmt.Println(req.GetItems())
	id, err := s.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &order_v1.CreateOrderResponse{
		Id:     id,
		Status: "Order created successfully",
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
		Id:      id,
		Message: "Changed",
	}, nil
}

func (s *Server) CancelOrder(ctx context.Context, req *order_v1.ChangeOrderStatusRequest) (*order_v1.ChangeOrderStatusResponse, error) {
	id, err := s.orderService.CancelOrder(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, app_errors.ErrCantCancelOrder) {
			return &order_v1.ChangeOrderStatusResponse{
				Id:      id,
				Message: "can't cancel order",
			}, nil
		}
		return nil, err
	}
	return &order_v1.ChangeOrderStatusResponse{
		Id:      id,
		Message: "order successfully cancelled",
	}, nil
}
