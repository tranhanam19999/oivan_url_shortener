package httpurlshortener

import (
	"errors"
	"net/http"

	domain "url-shortener/internal/service/urlshortener"

	"github.com/labstack/echo/v4"
)

func mapError(err error) *echo.HTTPError {
	switch {
	case errors.Is(err, domain.ErrURLNotFound):
		return echo.NewHTTPError(http.StatusNotFound, "URL not found")

	case errors.Is(err, domain.ErrInvalidURL):
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid URL provided")

	default:
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
}
