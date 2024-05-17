package main

import (
	"google.golang.org/grpc"
	"kitchenService/internal/config"
	"kitchenService/internal/repositories"
	"kitchenService/internal/server"
	"kitchenService/internal/services"
	kitchen_v1 "kitchenService/proto/v1"
	"log"
	"net"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	kitchenRepository, err := repositories.NewKitchenRepository(cfg)
	if err != nil {
		log.Fatalf("failed to create order repository: %v", err)
	}
	kitchenService := services.NewKitchenService(kitchenRepository)
	srv := server.NewServer(kitchenService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	kitchen_v1.RegisterKitchenServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
