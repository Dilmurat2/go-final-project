package services

import (
	"context"
	"menu/internal/models"
	"menu/internal/ports"
	"time"
)

type menuService struct {
	menuRepo ports.MenuRepository
}

func NewMenuService(menuRepo ports.MenuRepository) ports.MenuService {
	return &menuService{
		menuRepo: menuRepo,
	}
}

func (m menuService) GetAll(ctx context.Context) ([]*models.Menu, error) {

	return m.menuRepo.GetAll(ctx)
}

func (m menuService) GetByID(ctx context.Context, id string) (*models.Menu, error) {
	return m.menuRepo.GetByID(ctx, id)
}

func (m menuService) Create(ctx context.Context, menu *models.Menu) (*models.Menu, error) {
	newMenu := models.NewMenu(menu.Name, menu.Description)
	return m.menuRepo.Create(ctx, newMenu)
}

func (m menuService) Update(ctx context.Context, menu *models.Menu) (*models.Menu, error) {
	menu.UpdatedAt = time.Now().String()
	return m.menuRepo.Update(ctx, menu)
}

func (m menuService) Delete(ctx context.Context, id string) error {
	return m.menuRepo.Delete(ctx, id)
}
