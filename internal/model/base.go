package model

import "time"

type Base struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (b *Base) BeforeCreate() {
	now := time.Now()

	if b.CreatedAt == nil {
		b.CreatedAt = &now
	}

	if b.UpdatedAt == nil {
		b.UpdatedAt = &now
	}
}

func (b *Base) BeforeUpdate() {
	// Set UpdatedAt timestamp
	now := time.Now()
	b.UpdatedAt = &now
}
