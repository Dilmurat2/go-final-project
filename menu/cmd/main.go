package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"menu/internal/config"
	"menu/internal/repositories"
	"menu/internal/server"
	"menu/internal/services"
	"menu/proto/menu"
	"net"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	menuRepository, err := repositories.NewMenuRepository(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	menuService := services.NewMenuService(menuRepository)

	srv := server.NewServer(menuService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", "50052"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	menu.RegisterMenuServiceServer(s, srv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
