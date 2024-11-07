package error_util

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CustomErrorType string

const (
	VALIDATION_ERROR      CustomErrorType = "VALIDATION ERROR"
	BAD_REQUEST_ERROR     CustomErrorType = "BAD REQUEST"
	INTERNAL_SERVER_ERROR CustomErrorType = "INTERNAL SERVER ERROR"
)

func (c CustomErrorType) String() string {
	return string(c)
}

type CustomError struct {
	errorType CustomErrorType

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

func (e *CustomError) SetErrorType(errorType CustomErrorType) *CustomError {
	e.errorType = errorType
	return e
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

func (e *CustomError) IsValidationError() bool {
	return e.errorType == VALIDATION_ERROR
}

/*Errors factory*/
func ErrorValidation(errValidation fiber.Map) *CustomError {
	return NewCustomError().
		SetCode(http.StatusBadRequest).
		SetErrorType(VALIDATION_ERROR).
		SetMessage(VALIDATION_ERROR.String()).
		SetFields(errValidation)
}
func Error400(msg string) *CustomError {
	return NewCustomError().
		SetCode(http.StatusBadRequest).
		SetErrorType(BAD_REQUEST_ERROR).
		SetMessage(msg)
}
func Error500(msg string) *CustomError {
	return NewCustomError().
		SetCode(http.StatusInternalServerError).
		SetErrorType(INTERNAL_SERVER_ERROR).
		SetMessage(msg)
}

/*Errors checking*/
func IsCustomError(err error) (customError *CustomError, isCustomError bool) {
	ok := errors.As(err, &customError)
	return customError, ok
}
func IsRealError(err error) bool {
	if err != nil {
		customError, isCustomError := IsCustomError(err)
		if isCustomError {
			return !customError.IsValidationError() // condition may be updated
		}
	}
	return true
}
