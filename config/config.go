package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Config - server configuration structure
type Config struct {
	Env  string `envconfig:"ENV" default:"development"`
	Port int    `envconfig:"PORT" default:"8080"`

	DBHost string `envconfig:"DB_HOST"`
	DBPort int    `envconfig:"DB_PORT"`
	DBName string `envconfig:"DB_NAME"`
	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASS"`
}

type serviceCfg struct {
}

// New - Get server configurations
func New() *Config {
	var cfg Config

	// Process config.
	err := envconfig.Process("", &cfg)
	if err != nil {
		zap.S().Fatalf("Error decoding environment variables: %v", err)
		return nil
	}

	return &cfg
}
