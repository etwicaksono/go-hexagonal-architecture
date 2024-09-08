package error_util

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	VALIDATION_ERROR = "VALIDATION ERROR"
	INVALID_REQUEST  = "INVALID REQUEST"
	INTERNAL_SERVER  = "INTERNAL SERVER ERROR"
	NOT_FOUND        = "NOT FOUND ERROR"
	UNAUTHORIZED     = "UNAUTHORIZED"
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

func IsCustomError(err error) (customError *CustomError, isCustomError bool) {
	ok := errors.As(err, &customError)
	return customError, ok
}

func ValidationError(errValidation fiber.Map) *CustomError {
	return NewCustomError().
		SetCode(http.StatusBadRequest).
		SetMessage(VALIDATION_ERROR).
		SetFields(errValidation)
}
