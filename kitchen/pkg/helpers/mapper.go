package helpers

import (
	"kitchenService/internal/models"
	pbkitchen "kitchenService/proto/v1"
)

func OrderModelToProtobuf(order *models.Order) *pbkitchen.Order {
	var pbOrder pbkitchen.Order

	pbOrder.Id = order.ID
	pbOrder.Status = string(order.Status)
	for _, item := range order.Items {
		pbOrder.Items = append(pbOrder.Items, &pbkitchen.Item{
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &pbOrder
}

func OrderPbToModel(pbOrder *pbkitchen.Order) *models.Order {
	var order models.Order
	order.Status = models.OrderStatus(pbOrder.Status)
	order.ID = pbOrder.GetId()
	for _, item := range pbOrder.Items {
		order.Items = append(order.Items, models.Item{
			Name:  item.GetName(),
			Price: item.GetPrice(),
		})
	}
	return &order
}
