package helpers

import (
	"kitchenService/internal/models"
	order2 "kitchenService/proto/order"
)

func OrderModelToProtobuf(order *models.Order) *order2.Order {
	var pbOrder order2.Order

	pbOrder.Id = order.ID
	pbOrder.Status = string(order.Status)
	for _, item := range order.Items {
		pbOrder.Items = append(pbOrder.Items, &order2.Item{
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &pbOrder
}

func OrderPbToModel(pbOrder *order2.Order) *models.Order {
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
