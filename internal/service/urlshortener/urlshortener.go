package urlshortener

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"url-shortener/internal/dto"
	repository "url-shortener/internal/repository/urlshortener"
	"url-shortener/tools/utils"

	"gorm.io/gorm"
)

func (s *service) EncodeUrl(ctx context.Context, input dto.EncodeURLReq) (*dto.EncodeURLResp, error) {
	rec, err := s.repo.FindOne(repository.FindOneInput{
		OriginalURL: input.URL,
	})

	fmt.Println("err encoding ", err)
	fmt.Println("really?? ", errors.Is(err, gorm.ErrRecordNotFound))
	fmt.Println("condition ", err != nil && !errors.Is(err, gorm.ErrRecordNotFound))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("huh?")
		return nil, err
	}

	fmt.Println("Rec ", rec)
	if rec != nil {
		return &dto.EncodeURLResp{
			URL: rec.ShortURL,
		}, nil
	}

	fmt.Println("continues")
	fmt.Println("Ctx? ", ctx)
	id, err := s.repo.GetNextID(ctx)
	if err != nil {
		return nil, err
	}

	encodedID := utils.EncodeBase62(id)

	shortURL, err := url.JoinPath(s.baseShortenURL, encodedID)
	if err != nil {
		return nil, err
	}

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
