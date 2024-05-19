package repository

import "go-final-project/restaurant_service/model"

var restaurants = []model.Restaurant{
	{ID: 1, Name: "Restaurant 1", Address: "Address 1", Phone: "111-111-1111"},
	{ID: 2, Name: "Restaurant 2", Address: "Address 2", Phone: "222-222-2222"},
}

func GetRestaurants() []model.Restaurant {
	return restaurants
}

func CreateRestaurant(restaurant model.Restaurant) {
	restaurant.ID = len(restaurants) + 1
	restaurants = append(restaurants, restaurant)
}
