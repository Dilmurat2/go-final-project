package main

import (
	"go-final-project/restaurant_service/config"
	"go-final-project/restaurant_service/handler"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	http.HandleFunc("/restaurants", handler.GetRestaurantsHandler)
	http.HandleFunc("/restaurant", handler.CreateRestaurantHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
