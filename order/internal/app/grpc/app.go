package grpcapp

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
	"orderService/internal/ports"
	"orderService/internal/server"
	order_v1 "orderService/proto/order"
	"runtime/debug"
)

type App struct {
	gRPCServer *grpc.Server
}

func New(
	orderService ports.OrderService,
) *App {
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			slog.Error("Recovered from panic", slog.Any("panic", p), slog.String("stack", string(debug.Stack())))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
	))

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

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", "50051"))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	slog.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	slog.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", 50051))

	a.gRPCServer.GracefulStop()
}
