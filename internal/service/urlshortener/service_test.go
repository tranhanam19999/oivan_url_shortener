package urlshortener_test

import (
	"context"
	"errors"
	"testing"

	"url-shortener/internal/dto"
	"url-shortener/internal/model"
	repository "url-shortener/internal/repository/urlshortener"
	repoMocks "url-shortener/internal/repository/urlshortener/mocks"
	"url-shortener/internal/service/urlshortener"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestService_EncodeUrl_ReturnExisting(t *testing.T) {
	mockRepo := repoMocks.NewMockURLShortenerRepository(t)
	svc := urlshortener.NewService(mockRepo, "http://base")

	mockRepo.
		On("FindOne", repository.FindOneInput{
			OriginalURL: "https://example.com",
		}).
		Return(&model.URLShortener{
			ShortURL: "http://base/cb",
		}, nil)

	resp, err := svc.EncodeUrl(context.Background(), dto.EncodeURLReq{
		URL: "https://example.com",
	})
	require.NoError(t, err)
	require.Equal(t, "http://base/cb", resp.URL)

	mockRepo.AssertExpectations(t)
}

func TestService_EncodeUrl_CreateNew(t *testing.T) {
	mockRepo := new(repoMocks.MockURLShortenerRepository)
	svc := urlshortener.NewService(mockRepo, "http://base")

	mockRepo.
		On("FindOne", repository.FindOneInput{
			OriginalURL: "https://example.com",
		}).
		Return(nil, gorm.ErrRecordNotFound)

	mockRepo.
		On("GetNextID", mock.Anything).
		Return(int64(125), nil)

	mockRepo.
		On("UpdateMapping", repository.UpdateMappingInput{
			ID:           125,
			OriginalURL:  "https://example.com",
			ShortenedURL: "http://base/21",
		}).
		Return(nil)

	resp, err := svc.EncodeUrl(context.Background(), dto.EncodeURLReq{
		URL: "https://example.com",
	})
	require.NoError(t, err)
	require.Equal(t, "http://base/21", resp.URL)

	mockRepo.AssertExpectations(t)
}

func TestService_DecodeUrl_Success(t *testing.T) {
	mockRepo := new(repoMocks.MockURLShortenerRepository)
	svc := urlshortener.NewService(mockRepo, "http://base")

	mockRepo.
		On("FindOne", repository.FindOneInput{
			ShortenedURL: "http://base/cb",
		}).
		Return(&model.URLShortener{
			OriginalURL: "https://example.com",
		}, nil)

	resp, err := svc.DecodeUrl(context.Background(), dto.DecodeURLReq{
		URL: "http://base/cb",
	})
	require.NoError(t, err)
	require.Equal(t, "https://example.com", resp.URL)

	mockRepo.AssertExpectations(t)
}

func TestService_DecodeUrl_NotFound(t *testing.T) {
	mockRepo := new(repoMocks.MockURLShortenerRepository)
	svc := urlshortener.NewService(mockRepo, "http://base")

	mockRepo.
		On("FindOne", repository.FindOneInput{
			ShortenedURL: "http://base/missing",
		}).
		Return(nil, gorm.ErrRecordNotFound)

	resp, err := svc.DecodeUrl(context.Background(), dto.DecodeURLReq{
		URL: "http://base/missing",
	})
	require.Nil(t, resp)
	require.True(t, errors.Is(err, gorm.ErrRecordNotFound))

	mockRepo.AssertExpectations(t)
}
