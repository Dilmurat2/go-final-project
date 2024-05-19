package helpers

import (
	"menu/internal/models"
	menu2 "menu/proto/menu"
	"menu/proto/order"
)

func ItemProtoToModel(item *order.Item) *models.Item {
	return &models.Item{
		ID:          item.GetId(),
		Name:        item.GetName(),
		Price:       item.GetPrice(),
		Description: item.GetDescription(),
	}
}

func MenuModelToProto(menu *models.Menu) *menu2.Menu {
	pbMenu := &menu2.Menu{
		Id:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
	}

	for _, item := range menu.Items {
		pbMenu.Items = append(pbMenu.Items, &order.Item{
			Id:          item.ID,
			Name:        item.Name,
			Price:       item.Price,
			Description: item.Description,
		})
	}

	return pbMenu
}
