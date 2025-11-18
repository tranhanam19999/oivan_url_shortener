package urlshortener

import (
	"context"
	"errors"
	"fmt"
	"url-shortener/internal/dto"
	repository "url-shortener/internal/repository/urlshortener"
	"url-shortener/tools/utils"

	"gorm.io/gorm"
)

func (s *service) EncodeUrl(ctx context.Context, input dto.EncodeURLReq) (*dto.EncodeURLResp, error) {
	rec, err := s.repo.FindOne(repository.FindOneInput{
		OriginalURL: input.URL,
	})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if rec != nil {
		return &dto.EncodeURLResp{
			URL: rec.ShortURL,
		}, nil
	}

	id, err := s.repo.GetNextID(ctx)
	if err != nil {
		return nil, err
	}

	encodedID := utils.EncodeBase62(id)

	// url.JoinPath won't work on ips
	// shortURL, err := url.JoinPath(s.baseShortenURL, encodedID)
	shortURL := fmt.Sprintf("%s/%s", s.baseShortenURL, encodedID)

	err = s.repo.UpdateMapping(repository.UpdateMappingInput{
		ID:           id,
		OriginalURL:  input.URL,
		ShortenedURL: shortURL,
	})
	if err != nil {
		return nil, err
	}

	return &dto.EncodeURLResp{
		URL: shortURL,
	}, nil
}

func (s *service) DecodeUrl(ctx context.Context, input dto.DecodeURLReq) (*dto.DecodeURLResp, error) {
	rec, err := s.repo.FindOne(repository.FindOneInput{
		ShortenedURL: input.URL,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &dto.DecodeURLResp{
		URL: rec.OriginalURL,
	}, nil
}
