package http

import (
	httpCommon "github.com/aryahmph/url-shortener/common/http"
	httpLinkCommon "github.com/aryahmph/url-shortener/common/http/link"
	linkModel "github.com/aryahmph/url-shortener/internal/model/link"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d HTTPLinkDelivery) addLink(c echo.Context) error {
	ctx := c.Request().Context()

	link := &httpLinkCommon.AddLink{}
	if err := c.Bind(link); err != nil {
		return err
	}
	if err := c.Validate(link); err != nil {
		return err
	}

	id, err := d.linkUsecase.CreateLink(ctx, linkModel.Link{
		OriginalURL: link.URL,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, httpCommon.Response{
		Data: httpLinkCommon.Link{
			ID:          id,
			OriginalURL: link.URL,
		},
	})
}
