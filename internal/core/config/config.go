package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is the configuration for the application
type Config struct {
	Port      int `envconfig:"PORT" default:"8080"`
	DebugPort int `envconfig:"DEBUG_PORT" default:"8081"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
