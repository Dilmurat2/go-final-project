package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"orderService/config"
	"orderService/internal/repositories"
	"orderService/internal/server"
	"orderService/internal/services"
	order_v1 "orderService/proto/v1"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	orderRepository, err := repositories.NewOrderRepository(cfg)
	if err != nil {
		log.Fatalf("failed to create order repository: %v", err)
	}
	orderService := services.NewOrderService(orderRepository)
	srv := server.NewServer(orderService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	order_v1.RegisterOrderServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
