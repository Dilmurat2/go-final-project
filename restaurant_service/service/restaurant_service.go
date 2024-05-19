package service

import (
	"go-final-project/restaurant_service/model"
	"go-final-project/restaurant_service/repository"
)

func GetRestaurants() []model.Restaurant {
	return repository.GetRestaurants()
}

func CreateRestaurant(restaurant model.Restaurant) {
	repository.CreateRestaurant(restaurant)
}
