package errors

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrNoData                   = error_util.Error400("No data in result")
	ErrFailedToHashPassword     = error_util.Error500("Failed to hash password")
	ErrEmailAlreadyUsed         = error_util.Error400("Email already used")
	ErrUsernameAlreadyUsed      = error_util.Error400("Username already used")
	ErrEmailMustBeInValidFormat = error_util.Error400("Email must be in valid format")
	ErrUserNotFound             = error_util.Error400("User not found")
	ErrInternalServer           = error_util.Error500("Internal server error")
	ErrLoginCredentialInvalid   = error_util.Error401WithField("Invalid login credentials", fiber.Map{
		"username": "Invalid username or password",
		"password": "Invalid username or password",
	})
	ErrUnauthorized            = error_util.Error401("Unauthorized")
	ErrTokenInvalid            = error_util.Error400("Invalid token")
	ErrTokenClaimsParseFailed  = error_util.Error400("Could not parse token claims")
	ErrUnsupportedDbProtocol   = error_util.Error500("Unsupported database protocol, supported protocol: [mongo, mysql]")
	ErrInvalidLogLevel         = error_util.Error500("Invalid log level, available options: [debug, info, warn, error]")
	ErrNoObjectToInsert        = error_util.Error400("No object to insert")
	ErrFailedToDeleteTempFiles = error_util.Error500("Failed to delete temp files")
	ErrMinioNotInitialized     = error_util.Error500("Minio not initialized")
	ErrInvalidContext          = error_util.Error500("Invalid context")
	ErrInvalidTempFiles        = error_util.Error500("Invalid temp files")
	ErrInvalidNewFilePath      = error_util.Error500("Invalid new file path")
)
