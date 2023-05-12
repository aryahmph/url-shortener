package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	E *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewHTTPServer() HTTPServer {
	e := echo.New()
	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}
	h := HTTPServer{E: e}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = h.errorHandler

	return h
}
