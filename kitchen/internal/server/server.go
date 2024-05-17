package server

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"kitchenService/internal/ports"
	kitchen_v1 "kitchenService/proto/v1"
)

type Server struct {
	kitchen_v1.UnimplementedKitchenServiceServer
	ports.KitchenService
}

func NewServer(kitchenService ports.KitchenService) *Server {
	return &Server{KitchenService: kitchenService}
}

func (s Server) ProcessOrder(ctx context.Context, order *kitchen_v1.Order) (*kitchen_v1.ProcessOrderResponse, error) {

}

func (s Server) ChangeOrderStatus(ctx context.Context, request *kitchen_v1.ChangeOrderStatusRequest) (*emptypb.Empty, error) {

}
