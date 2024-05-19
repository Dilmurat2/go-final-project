package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
}

type Config struct {
	Database Database
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("MenuService", cfg); err != nil {
		return nil, fmt.Errorf("error processing MenuService env: %w", err)
	}
	return cfg, nil
}
