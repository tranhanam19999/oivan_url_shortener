package urlshortener

import (
	"context"
	repository "url-shortener/internal/repository/urlshortener"
)

type Service interface {
	// Encode will encode an url into a shorterned version
	// With a link contains the server's domain
	EncodeUrl(ctx context.Context, url string) (string, error)

	// While decode will decode the shortened url made by the encoded one
	// If there is no such url, it will return an error
	DecodeUrl(ctx context.Context, url string) (string, error)
}

type service struct {
	repo *repository.URLShortenerRepository
}

func NewService(repo *repository.URLShortenerRepository) Service {
	return &service{
		repo: repo,
	}
}
