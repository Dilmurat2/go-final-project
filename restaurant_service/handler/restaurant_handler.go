package handler

import (
	"encoding/json"
	"go-final-project/restaurant_service/model"
	"go-final-project/restaurant_service/service"
	"net/http"
)

func GetRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	restaurants := service.GetRestaurants()
	json.NewEncoder(w).Encode(restaurants)
}

func CreateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	var restaurant model.Restaurant
	json.NewDecoder(r.Body).Decode(&restaurant)
	service.CreateRestaurant(restaurant)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurant)
}
