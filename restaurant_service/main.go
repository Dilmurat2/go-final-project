package main

import (
	"log"
	"net/http"
	"restaurant_service/config"
	"restaurant_service/handler"
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
