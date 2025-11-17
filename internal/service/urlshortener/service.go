package urlshortener

import (
	"context"
	"url-shortener/internal/dto"
	repository "url-shortener/internal/repository/urlshortener"
)

type Service interface {
	// Encode will encode an url into a shorterned version
	// With a link contains the server's domain
	EncodeUrl(ctx context.Context, input dto.EncodeURLReq) (*dto.EncodeURLResp, error)

	// While decode will decode the shortened url made by the encoded one
	// If there is no such url, it will return an error
	DecodeUrl(ctx context.Context, input dto.DecodeURLReq) (*dto.DecodeURLResp, error)
}

type service struct {
	repo           repository.URLShortenerRepository
	baseShortenURL string
}

func NewService(repo repository.URLShortenerRepository, baseShortenURL string) Service {
	return &service{
		repo:           repo,
		baseShortenURL: baseShortenURL,
	}
}
