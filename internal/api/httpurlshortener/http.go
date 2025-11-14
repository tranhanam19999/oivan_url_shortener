package httpurlshortener

import (
	"context"
	"net/http"
	"url-shortener/internal/dto"
	"url-shortener/tools/utils"

	"github.com/labstack/echo/v4"
)

type Service interface {
	EncodeUrl(ctx context.Context, input dto.EncodeURLReq) (*dto.EncodeURLResp, error)
	DecodeUrl(ctx context.Context, input dto.DecodeURLReq) (*dto.DecodeURLResp, error)
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
	r := EncodeURLInput{}
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if valid := utils.IsValidUrl(r.URL); !valid {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid url format")
	}

	resp, err := h.svc.EncodeUrl(c.Request().Context(), dto.EncodeURLReq{
		URL: r.URL,
	})

	if err != nil {
		httpErr := mapError(err)
		return c.JSON(httpErr.Code, httpErr.Message)
	}

	return c.JSON(200, resp)
}

func (h *HTTP) decodeUrlHandler(c echo.Context) error {
	r := DecodeURLInput{}
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if valid := utils.IsValidUrl(r.URL); !valid {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid url format")
	}

	resp, err := h.svc.DecodeUrl(c.Request().Context(), dto.DecodeURLReq{
		URL: r.URL,
	})

	if err != nil {
		httpErr := mapError(err)
		return c.JSON(httpErr.Code, httpErr.Message)
	}

	return c.JSON(200, resp)
}
