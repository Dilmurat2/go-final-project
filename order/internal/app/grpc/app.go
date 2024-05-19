package grpcapp

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"orderService/internal/ports"
	"orderService/internal/server"
	order_v1 "orderService/proto/order"
)

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Логирование метаданных
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	} else {
		log.Println("No metadata received")
	}

	// Продолжение выполнения запроса
	return handler(ctx, req)
}

type App struct {
	gRPCServer *grpc.Server
}

func New(orderService ports.OrderService) *App {
	// Создаем цепочку перехватчиков с нашим middleware логирования
	chain := grpc.UnaryInterceptor(loggingInterceptor)

	gRPCServer := grpc.NewServer(chain)

	srv := server.NewServer(orderService)
	order_v1.RegisterOrderServiceServer(gRPCServer, srv)

	return &App{
		gRPCServer: gRPCServer,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", "12201"))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Println("gRPC server started", "addr", l.Addr().String())
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	log.Println("stopping gRPC server", "port", 50051)
	a.gRPCServer.GracefulStop()
}
