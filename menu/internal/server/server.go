package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"menu/internal/models"
	"menu/internal/ports"
	menu_v1 "menu/proto/v1"
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
			Id:   menu.ID,
			Name: menu.Name,
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
