package usrepository

import (
	"gorm.io/gorm"
)

type URLShortenerRepository interface {
	Save(originalURL, shortURL string) error
	FindOriginalURL(url string) (string, error)
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) URLShortenerRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Save(originalURL, shortURL string) error {
	// Implementation to save the URL mapping in the database
	return nil
}

func (r *repo) FindOriginalURL(shortURL string) (string, error) {
	// Implementation to find the original URL from the shortened URL in the database
	return "", nil
}
