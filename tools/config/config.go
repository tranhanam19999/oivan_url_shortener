package config

import (
	"log"
	"url-shortener/tools/utils"

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

	cfg.App.SBaseURL = utils.BuildShortenUrlFromConfig(utils.BuildUrlFromConfigInput{
		Host:  cfg.App.Host,
		Port:  cfg.App.Port,
		Stage: cfg.Stage,
	})

	return &cfg
}

// This is only for intergration testing
func LoadTest() *Config {
	_ = godotenv.Load("../../.env.test") // load .env into environment
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}

	cfg.App.SBaseURL = utils.BuildShortenUrlFromConfig(utils.BuildUrlFromConfigInput{
		Host:  cfg.App.Host,
		Port:  cfg.App.Port,
		Stage: cfg.Stage,
	})

	return &cfg
}
