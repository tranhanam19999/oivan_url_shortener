package config

import (
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

	return &cfg
}
