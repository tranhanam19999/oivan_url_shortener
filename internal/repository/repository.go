package repository

import (
	usrepository "url-shortener/internal/repository/urlshortener"

	"gorm.io/gorm"
)

type Repository struct {
	UrlShortener usrepository.URLShortenerRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		UrlShortener: usrepository.New(db),
	}
}
