package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type database struct {
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
}

type Config struct {
	Database              database
	KitchenServiceAddress string `envconfig:"KITCHEN_SERVICE_ADDRESS"`
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
