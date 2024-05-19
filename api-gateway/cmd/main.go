package main

import (
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

		// Добавляем заголовки метаданных Postman
		r.Header.Set("user-agent", "PostmanRuntime/7.39.0")
		r.Header.Set(":authority", "localhost:12201")

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading gRPC Gateway request body: %v", err)
		} else {
			log.Printf("JSON Request: %s", buf)
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}

	loggingMux := LoggingMiddleware(mux)

	log.Printf("server listening at :8081")
	if err := http.ListenAndServe(":8081", loggingMux); err != nil {
		panic(err)
	}
}
