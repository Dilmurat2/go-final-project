package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"menu/internal/models"
	"menu/internal/ports"
	menu_v1 "menu/proto/menu"
	"menu/proto/order"
	"time"
)

type Server struct {
	menu_v1.UnimplementedMenuServiceServer
	menuService ports.MenuService
}

func NewServer(menuService ports.MenuService) *Server {
	return &Server{
		menuService: menuService,
	}
}

func (s *Server) GetAllMenus(ctx context.Context, req *emptypb.Empty) (*menu_v1.GetMenuResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*250)
	defer cancel()

	menus, err := s.menuService.GetAll(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("service timed out after 1 second")
		}
		return nil, err
	}

	var pbMenus []*menu_v1.Menu
	for _, menu := range menus {
		pbMenus = append(pbMenus, &menu_v1.Menu{
			Id:          menu.ID,
			Name:        menu.Name,
			Description: menu.Description,
		})
	}

	return &menu_v1.GetMenuResponse{
		Menus: pbMenus,
	}, nil
}

func (s *Server) CreateMenu(ctx context.Context, req *menu_v1.Menu) (*menu_v1.Menu, error) {
	menu := &models.Menu{
		Name:        req.Name,
		Description: req.Description,
	}

	newMenu, err := s.menuService.Create(ctx, menu)
	if err != nil {
		return nil, err
	}

	return &menu_v1.Menu{
		Id:          newMenu.ID,
		Name:        newMenu.Name,
		Description: newMenu.Description,
	}, nil
}

func (s *Server) UpdateMenu(ctx context.Context, req *menu_v1.Menu) (*menu_v1.Menu, error) {
	menu := &models.Menu{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}

	for _, item := range req.Items {
		menu.Items = append(menu.Items, models.Item{
			Name:  item.Name,
			Price: item.Price,
		})
	}

	updatedMenu, err := s.menuService.Update(ctx, menu)
	if err != nil {
		return nil, err
	}

	return &menu_v1.Menu{
		Id:          updatedMenu.ID,
		Name:        updatedMenu.Name,
		Description: updatedMenu.Description,
	}, nil
}

func (s *Server) GetMenu(ctx context.Context, req *menu_v1.GetMenuRequest) (*menu_v1.Menu, error) {
	menu, err := s.menuService.GetByID(ctx, req.GetMenuId())
	if err != nil {
		return nil, err
	}

	pbMenu := &menu_v1.Menu{
		Id:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
	}

	for _, item := range menu.Items {
		pbMenu.Items = append(pbMenu.Items, &order.Item{
			Id:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return pbMenu, nil
}
