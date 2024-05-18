package main

import (
	"api-gateway/config"
	"api-gateway/proto/kitchen"
	"api-gateway/proto/menu"
	"api-gateway/proto/order"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, cfg.OrderServiceAddr, opts); err != nil {
		log.Fatalf("failed to register OrderService handler: %v", err)
	}

	if err := menu.RegisterMenuServiceHandlerFromEndpoint(ctx, mux, cfg.MenuServiceAddr, opts); err != nil {
		log.Fatalf("failed to register MenuService handler: %v", err)
	}

	if err := kitchen.RegisterKitchenServiceHandlerFromEndpoint(ctx, mux, cfg.KitchenServiceAddr, opts); err != nil {
		log.Fatalf("failed to register KitchenService handler: %v", err)
	}

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
