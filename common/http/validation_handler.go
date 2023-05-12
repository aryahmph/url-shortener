package http

import (
	"errors"
	errorCommon "github.com/aryahmph/url-shortener/common/error"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		var ves validator.ValidationErrors
		errors.As(err, &ves)
		keys := make(map[string]string)
		for _, ve := range ves {
			keys[strings.ToLower(ve.Field())] = ve.Tag()
		}

		return &echo.HTTPError{
			Code: http.StatusUnprocessableEntity,
			Message: errorCommon.ClientError{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    ves.Error(),
				Errors:     keys,
			},
		}
	}
	return nil
}
