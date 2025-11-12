package httpurlshortener

import (
	"context"
	"url-shortener/internal/dto"

	"github.com/labstack/echo/v4"
)

type Service interface {
	EncodeUrl(ctx context.Context, url string) (string, error)
	DecodeUrl(ctx context.Context, url string) (string, error)
}

type HTTP struct {
	svc Service
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{
		svc: svc,
	}

	// swagger:route POST /url-shortener/encode url-shortener encodeUrl
	//
	// Encode a given URL to a shortened version.
	//
	// Responses:
	//   200: encodeResponse
	//   400: badRequestResponse
	//   500: internalServerErrorResponse
	eg.POST("/encode", h.encodeUrlHandler)

	// swagger:route POST /url-shortener/decode url-shortener decodeUrl
	//
	// Decode a shortened URL to its original form.
	//
	// Responses:
	//   200: decodeResponse
	//   400: badRequestResponse
	//   404: notFoundResponse
	//   500: internalServerErrorResponse
	eg.POST("/decode", h.decodeUrlHandler)
}

func (h *HTTP) encodeUrlHandler(c echo.Context) error {
	r := dto.EncodeURLReq{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	h.svc.EncodeUrl(c.Request().Context(), r.URL)
	return nil
}

func (h *HTTP) decodeUrlHandler(c echo.Context) error {
	return nil
}
