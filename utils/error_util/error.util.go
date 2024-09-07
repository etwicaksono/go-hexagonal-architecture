package error_util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
)

type CustomError struct {
	Code    int
	Message string
	Fields  fiber.Map
}

func NewCustomError() *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
}

func IsCustomError(err error) (customError *CustomError, isCustomError bool) {
	ok := errors.As(err, &customError)
	return customError, ok
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) SetCode(code int) *CustomError {
	e.Code = code
	return e
}

func (e *CustomError) SetMessage(msg string) *CustomError {
	e.Message = msg
	return e
}

func (e *CustomError) SetFields(fields fiber.Map) *CustomError {
	e.Fields = fields
	return e
}
