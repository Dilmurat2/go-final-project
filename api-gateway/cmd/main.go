package main

import (
	"api-gateway/config"
	"api-gateway/proto/kitchen"
	"api-gateway/proto/menu"
	"api-gateway/proto/order"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.RequestURI)
		log.Printf("Headers: %v", r.Header)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
		} else {
			log.Printf("Body: %s", body)
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, cfg.OrderServiceAddr, opts)
	if err != nil {
		panic(err)
	}

	err = kitchen.RegisterKitchenServiceHandlerFromEndpoint(ctx, mux, cfg.KitchenServiceAddr, opts)
	if err != nil {
		panic(err)
	}

	err = menu.RegisterMenuServiceHandlerFromEndpoint(ctx, mux, cfg.MenuServiceAddr, opts)
	if err != nil {
		panic(err)
	}

	loggingMux := LoggingMiddleware(mux)

	log.Printf("server listening at :8081")
	if err := http.ListenAndServe(":8080", loggingMux); err != nil {
		panic(err)
	}
}
