package httpredirect

import (
	"context"
	"net/http"
	"url-shortener/internal/dto"
	"url-shortener/tools/utils"

	"github.com/labstack/echo/v4"
)

type Service interface{}

type USService interface {
	DecodeUrl(ctx context.Context, input dto.DecodeURLReq) (*dto.DecodeURLResp, error)
}

type HTTP struct {
	svc   Service
	usSvc USService
}

func NewHTTP(svc Service, usSvc USService, eg *echo.Group) {
	h := HTTP{
		svc:   svc,
		usSvc: usSvc,
	}

	// swagger:route POST /url-shortener/decode url-shortener decodeUrl
	//
	// Decode a shortened URL to its original form.
	//
	// Responses:
	//   201: redirect
	//   400: badRequestResponse
	//   404: notFoundResponse
	//   500: internalServerErrorResponse
	eg.GET("/:code", h.redirectUrlHandler)
}

func (h *HTTP) redirectUrlHandler(c echo.Context) error {
	r := RedirectUrlInput{}
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	reqUrl := utils.BuildUrlFromRequest(c.Request())
	if valid := utils.IsValidUrl(reqUrl); !valid {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid url")
	}

	resp, err := h.usSvc.DecodeUrl(c.Request().Context(), dto.DecodeURLReq{
		URL: reqUrl,
	})

	if err != nil {
		httpErr := mapError(err)
		return c.JSON(httpErr.Code, httpErr.Message)
	}

	return c.Redirect(http.StatusFound, resp.URL)
}
