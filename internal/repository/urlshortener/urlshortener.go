package repository

import (
	"url-shortener/internal/model"

	"gorm.io/gorm"
)

type URLShortenerRepository interface {
	// Create will do insert a new US record
	Create(input CreateInput) error
	// FindOne will find a URLShortener by original URL
	FindOne(input FindOneInput) (*model.URLShortener, error)
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) URLShortenerRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(input CreateInput) error {
	if err := r.db.Create(&model.URLShortener{
		OriginalURL: input.OriginalURL,
		ShortURL:    input.ShortenedURL,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) FindOne(input FindOneInput) (*model.URLShortener, error) {
	var us *model.URLShortener
	if err := r.db.Where("original_url = ?", input.URL).First(&us).Error; err != nil {
		return nil, err
	}

	return us, nil
}
