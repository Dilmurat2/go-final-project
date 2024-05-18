package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"menu/proto/menu"
	"time"
)

type Menu struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Items       []Item `json:"items" bson:"items"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	UpdatedAt   string `json:"updated_at" bson:"updated_at"`
	DeletedAt   string `json:"deleted_at" bson:"deleted_at"`
	IsActive    bool   `json:"is_active" bson:"is_active"`
}

func NewMenu(name, description string) *Menu {
	return &Menu{
		ID:          primitive.NewObjectID().Hex(),
		Name:        name,
		Description: description,
		Items:       make([]Item, 0),
		IsActive:    true,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
}

func PbMenuToModel(pbMenu *menu.Menu) *Menu {
	menu := &Menu{
		ID:          pbMenu.Id,
		Name:        pbMenu.Name,
		Description: pbMenu.Description,
	}

	for _, pbItem := range pbMenu.Items {
		item := Item{
			ID:    pbItem.Id,
			Name:  pbItem.Name,
			Price: pbItem.Price,
		}
		menu.Items = append(menu.Items, item)
	}

	return menu
}
