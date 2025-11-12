package dbutil

import (
	"fmt"
	"url-shortener/tools/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.DBConfig) (*gorm.DB, error) {
	dsn := openConnection(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func openConnection(cfg *config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)
}
