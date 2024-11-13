package rest

import (
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/middleware"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	utils2 "github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2/utils"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func NewRestApp(
	cfg config.Config,
	route Router,
) *fiber.App {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Use absolute path for the templates directory
	templatePath := filepath.Join(wd, "/docs")
	engine := html.New(templatePath, ".gohtml")

	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  cfg.App.IdleTimeout,
		WriteTimeout: cfg.App.WriteTimeout,
		ReadTimeout:  cfg.App.ReadTimeout,
		Prefork:      cfg.App.RestPrefork,
		Views:        engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			status := fiber.ErrInternalServerError.Message
			message := err.Error()
			errorMap := map[string]any{}

			// Retrieve the custom status code if it's a *fiber.Error
			var fiberError *fiber.Error
			if errors.As(err, &fiberError) {
				code = fiberError.Code
				status = utils.StatusMessage(fiberError.Code)
				message = fiberError.Error()
			}

			// Retrieve the custom status code if it's an utils2.CustomError
			var customError *utils2.CustomError
			if errors.As(err, &customError) {
				code = customError.Code
				status = utils.StatusMessage(customError.Code)
				errorMap = customError.Fields
				message = customError.Error()
			}

			return rest_util.ResponseGeneral(ctx, code, model.Response[any]{
				Code:    code,
				Status:  status,
				Message: message,
				Errors:  errorMap,
			})
		},
	})

	if cfg.App.RestRecovery {
		fiberApp.Use(recover.New(recover.Config{
			EnableStackTrace: cfg.App.RestEnableStackTrace,
		})) // Panic Handler
	}

	// Middleware before route
	middleware.CorsMiddleware(fiberApp, cfg.App)
	middleware.UnprocessableEntityMiddleware(fiberApp)

	// SetRoute
	SetRoute(fiberApp, route)

	// Static files
	docPath := filepath.Join(wd, "/docs")
	staticFiles := map[string]string{
		"/docs": docPath,
	}
	for key, value := range staticFiles {
		fiberApp.Static(key, value)
	}

	// Middleware after route
	middleware.NotFoundMiddleware(fiberApp)

	return fiberApp
}
