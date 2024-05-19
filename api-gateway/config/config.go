package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MenuServiceAddr    string `envconfig:"MENU_SERVICE_ADDR" default:"localhost:50051"`
	OrderServiceAddr   string `envconfig:"ORDER_SERVICE_ADDR" default:"localhost:50052"`
	KitchenServiceAddr string `envconfig:"KITCHEN_SERVICE_ADDR" default:"localhost:50053"`
}

func New() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("MenuService", cfg); err != nil {
		return nil, fmt.Errorf("error processing MenuService env: %w", err)
	}
	return cfg, nil
}
