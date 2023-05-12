package http

import (
	httpCommon "github.com/aryahmph/url-shortener/common/http"
	httpLinkCommon "github.com/aryahmph/url-shortener/common/http/link"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d HTTPLinkDelivery) getLink(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	link, err := d.linkUsecase.GetLink(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, httpCommon.Response{
		Data: httpLinkCommon.Link{
			ID:          id,
			OriginalURL: link.OriginalURL,
		},
	})
}
