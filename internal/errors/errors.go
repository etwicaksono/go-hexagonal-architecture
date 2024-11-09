package errors

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrNoData                   = error_util.Error400("no data in result")
	ErrFailedToHashPassword     = error_util.Error500("failed to hash password")
	ErrEmailAlreadyUsed         = error_util.Error400("email already used")
	ErrUsernameAlreadyUsed      = error_util.Error400("username already used")
	ErrEmailMustBeInValidFormat = error_util.Error400("email must be in valid format")
	ErrUserNotFound             = error_util.Error400("user not found")
	ErrInternalServer           = error_util.Error500("internal server error")
	ErrLoginCredentialInvalid   = error_util.Error401WithField("invalid login credentials", fiber.Map{
		"username": "Invalid username or password",
		"password": "Invalid username or password",
	})
	ErrUnauthorized = error_util.Error401("unauthorized")
	ErrTokenInvalid = error_util.Error400("invalid token")
)
