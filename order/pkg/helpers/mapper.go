package helpers

import (
	"fmt"
	"orderService/internal/models"
	order_v1 "orderService/proto/order"
)

func OrderProtobufToModel(pbOrder *order_v1.Order) *models.Order {
	var order models.Order
	order.Status = models.OrderStatus(pbOrder.GetStatus())
	order.ID = pbOrder.GetId()
	for _, item := range pbOrder.Items {
		fmt.Println(item.GetName())
		order.Items = append(order.Items, &models.Item{
			Name:  item.GetName(),
			Price: item.GetPrice(),
		})
	}
	return &order
}

func OrderModelToProtobuf(order *models.Order) *order_v1.Order {
	pbOrder := &order_v1.Order{
		Id:     order.ID,
		Status: string(order.Status),
	}
	for _, item := range order.Items {
		pbOrder.Items = append(pbOrder.Items, &order_v1.Item{
			Name:  item.Name,
			Price: item.Price,
		})
	}
	return pbOrder
}
