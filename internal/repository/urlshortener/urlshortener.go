package repository

import (
	"context"
	"fmt"
	"url-shortener/internal/model"

	"gorm.io/gorm"
)

type URLShortenerRepository interface {
	// This will update to an existing US record
	UpdateMapping(input UpdateMappingInput) error

	// FindOne will find a URLShortener by it's input fields param
	FindOne(input FindOneInput) (*model.URLShortener, error)

	GetNextID(ctx context.Context) (int64, error)
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) URLShortenerRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) UpdateMapping(input UpdateMappingInput) error {
	if err := r.db.Where("id = ?", input.ID).Omit("id").Updates(&model.URLShortener{
		OriginalURL: input.OriginalURL,
		ShortURL:    input.ShortenedURL,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) FindOne(input FindOneInput) (*model.URLShortener, error) {
	var us *model.URLShortener
	dbCtx := r.db.Model(&model.URLShortener{})

	if input.OriginalURL != "" {
		dbCtx = dbCtx.Where("original_url = ?", input.OriginalURL)
	}

	if input.ShortenedURL != "" {
		dbCtx = dbCtx.Where("short_url = ?", input.ShortenedURL)
	}

	if err := dbCtx.First(&us).Error; err != nil {
		return nil, err
	}

	return us, nil
}

func (r *repo) GetNextID(ctx context.Context) (int64, error) {
	var id *int64

	fmt.Println("&id ", &id)
	err := r.db.WithContext(ctx).Raw(`INSERT INTO url_shorteners (original_url) VALUES ('') RETURNING id`).Scan(&id).Error
	return *id, err
}
