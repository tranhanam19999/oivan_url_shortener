package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Load() *Config {
	// Reads from .env values
	_ = godotenv.Load() // load .env into environment
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}

	cfg.App.BaseURL = buildBaseURL(cfg)
	return &cfg
}

func buildBaseURL(cfg Config) string {
	if cfg.Stage == "local" {
		return fmt.Sprintf("http://%s:%s", cfg.App.Host, cfg.App.Port)
	}

	return cfg.App.Host
}
