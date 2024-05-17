package app

import (
	"log"
	"orderService/config"
	grpcapp "orderService/internal/app/grpc"
	"orderService/internal/repositories"
	"orderService/internal/services"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(cfg *config.Config) *App {
	orderRepository, err := repositories.NewOrderRepository(cfg)
	if err != nil {
		log.Fatalf("failed to create order repository: %v", err)
	}

	kc := services.NewKitchenProxy(cfg)
	orderService := services.NewOrderService(orderRepository, kc)

	grpcServer := grpcapp.New(orderService)
	return &App{GRPCServer: grpcServer}
}
