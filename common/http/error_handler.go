package http

import (
	errorCommon "github.com/aryahmph/url-shortener/common/error"
	"github.com/labstack/echo/v4"
)

func (h HTTPServer) errorHandler(err error, c echo.Context) {
	if he, ok := err.(*errorCommon.ClientError); ok {
		err = &echo.HTTPError{
			Code:    he.StatusCode,
			Message: he,
		}
	} else if he, ok = errorCommon.DomainErrorTranslatorDirectories[err]; ok {
		err = &echo.HTTPError{
			Code:    he.StatusCode,
			Message: he,
		}
	}

	h.E.DefaultHTTPErrorHandler(err, c)
}
