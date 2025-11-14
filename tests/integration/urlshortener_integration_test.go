package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/api/httpurlshortener"
	"url-shortener/internal/dto"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	"url-shortener/internal/service/urlshortener"
	"url-shortener/tools/config"
	"url-shortener/tools/dbutil"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeFlow(t *testing.T) {
	// Load config & in-memory DB
	cfg := config.LoadTest()
	db, err := dbutil.New(&cfg.DB)
	require.NoError(t, err)

	// Setup Echo
	e := echo.New()
	g := e.Group("/url-shortener")

	// Setup repository + service + HTTP handlers
	repos := repository.NewRepository(db)
	svc := urlshortener.NewService(repos.UrlShortener, cfg.App.BaseURL)
	httpurlshortener.NewHTTP(svc, g)

	// Auto migrate test DB
	db.AutoMigrate(&model.URLShortener{})

	// =====================================================
	// 1. Encode request
	// =====================================================
	encodeBody := `{"url":"https://google.com"}`
	req := httptest.NewRequest(http.MethodPost, "/url-shortener/encode", bytes.NewBufferString(encodeBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var encodeResp dto.EncodeURLResp
	err = json.Unmarshal(rec.Body.Bytes(), &encodeResp)
	require.NoError(t, err)
	require.NotEmpty(t, encodeResp.URL)

	// =====================================================
	// 2. Decode request
	// =====================================================
	decodeBody := fmt.Sprintf(`{"url": "%s"}`, encodeResp.URL)

	req2 := httptest.NewRequest(http.MethodPost, "/url-shortener/decode", bytes.NewBufferString(decodeBody))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec2 := httptest.NewRecorder()
	e.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusOK, rec2.Code)

	var decodeResp dto.DecodeURLResp
	err = json.Unmarshal(rec2.Body.Bytes(), &decodeResp)
	require.NoError(t, err)

	require.Equal(t, "https://google.com", decodeResp.URL)
}
