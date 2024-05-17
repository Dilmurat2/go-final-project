package main

import (
	"log"
	"log/slog"
	"orderService/config"
	"orderService/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	application := app.New(cfg)
	go func() {
		if err := application.GRPCServer.Run(); err != nil {
			log.Fatalf("failed to run grpc server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	slog.Info("Gracefully stopped")
}
