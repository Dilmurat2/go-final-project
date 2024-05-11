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
}
