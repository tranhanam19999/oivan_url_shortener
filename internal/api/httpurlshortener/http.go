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

	eg.POST("/encode", h.encodeUrlHandler)

	eg.POST("/decode", h.decodeUrlHandler)
}

// EncodeUrlHandler godoc
// @Summary Encode URL
// @Description Encode a given URL to a shortened version.
// @Tags url-shortener
// @Accept json
// @Produce json
// @Param input body httpurlshortener.EncodeURLInput true "URL Input to be encoded"
// @Success 200 {object} dto.EncodeURLResp
// @Failure 400 {object} dto.EchoHttpErrorResp
// @Failure 500 {object} dto.EchoHttpErrorResp
// @Router /url-shortener/encode [post]
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

// DecodeUrlHandler godoc
// @Summary Decode URL
// @Description Decode a given URL to a shortened version.
// @Tags url-shortener
// @Accept json
// @Produce json
// @Param input body httpurlshortener.DecodeURLInput true "URL Input to be decoded"
// @Success 200 {object} dto.DecodeURLResp
// @Failure 400 {object} dto.EchoHttpErrorResp
// @Failure 500 {object} dto.EchoHttpErrorResp
// @Router /url-shortener/decode [post]
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
