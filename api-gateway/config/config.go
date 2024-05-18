package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MenuServiceAddr    string `envconfig:"MENU_SERVICE_ADDR" default:"localhost:50051"`
	OrderServiceAddr   string `envconfig:"ORDER_SERVICE_ADDR" default:"localhost:50052"`
	KitchenServiceAddr string `envconfig:"KITCHEN_SERVICE_ADDR" default:"localhost:50053"`
}

func New() (*Config, error) {
	cfg := new(Config)
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %e", err)
	}
	if err := envconfig.Process("MenuService", cfg); err != nil {
		return nil, fmt.Errorf("error processing MenuService env: %w", err)
	}
	return cfg, nil
}
