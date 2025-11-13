package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	if b.CreatedAt == nil {
		b.CreatedAt = &now
	}

	if b.UpdatedAt == nil {
		b.UpdatedAt = &now
	}

	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	// Set UpdatedAt timestamp
	now := time.Now()
	b.UpdatedAt = &now

	return nil
}
