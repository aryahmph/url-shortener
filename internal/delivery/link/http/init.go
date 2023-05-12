package http

import (
	linkUCase "github.com/aryahmph/url-shortener/internal/usecase/link"
	"github.com/labstack/echo/v4"
)

type HTTPLinkDelivery struct {
	linkUsecase linkUCase.Usecase
}

func NewHTTPLinkDelivery(g *echo.Group, linkUsecase linkUCase.Usecase) HTTPLinkDelivery {
	h := HTTPLinkDelivery{linkUsecase: linkUsecase}

	g.POST("/links", h.addLink)
	g.GET("/links/:id", h.getLink)

	return h
}
