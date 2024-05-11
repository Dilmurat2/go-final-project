package ports

import (
	"context"
	"menu/internal/models"
)

type MenuRepository interface {
	GetAll(ctx context.Context) ([]*models.Menu, error)
	GetByID(ctx context.Context, id string) (*models.Menu, error)
	Create(ctx context.Context, menu *models.Menu) (*models.Menu, error)
	Update(ctx context.Context, menu *models.Menu) (*models.Menu, error)
	Delete(ctx context.Context, id string) error
	AddItem(ctx context.Context, menuID string, item *models.Item) (*models.Menu, error)
	DeleteItem(ctx context.Context, menuID string, itemID string) (*models.Menu, error)
}
