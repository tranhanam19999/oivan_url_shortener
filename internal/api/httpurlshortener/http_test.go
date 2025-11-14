package httpurlshortener

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/api/httpurlshortener/mocks"
	"url-shortener/internal/dto"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newEcho() *echo.Echo {
	e := echo.New()
	return e
}

func TestEncodeUrlHandler(t *testing.T) {
	mockSvc := new(mocks.MockService)

	e := newEcho()
	eg := e.Group("/url-shortener")
	NewHTTP(mockSvc, eg)

	body := `{"url":"https://example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/url-shortener/encode", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mockSvc.
		On("EncodeUrl", mock.Anything, dto.EncodeURLReq{
			URL: "https://example.com",
		}).
		Return(&dto.EncodeURLResp{URL: "http://base/cb"}, nil)

	// Let echo execute the route automatically
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	var resp dto.EncodeURLResp
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

	require.Equal(t, "http://base/cb", resp.URL)
	mockSvc.AssertExpectations(t)
}

func TestDecodeUrlHandler(t *testing.T) {
	mockSvc := new(mocks.MockService)

	e := newEcho()
	eg := e.Group("/url-shortener")
	NewHTTP(mockSvc, eg)

	body := `{"url":"http://base/cb"}`
	req := httptest.NewRequest(http.MethodPost, "/url-shortener/decode", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mockSvc.On("DecodeUrl", mock.Anything, dto.DecodeURLReq{
		URL: "http://base/cb",
	}).Return(&dto.DecodeURLResp{
		URL: "https://example.com",
	}, nil)

	e.ServeHTTP(rec, req)
	require.Equal(t, 200, rec.Code)

	var resp dto.DecodeURLResp
	json.Unmarshal(rec.Body.Bytes(), &resp)

	require.Equal(t, "https://example.com", resp.URL)
	mockSvc.AssertExpectations(t)
}
