package rest

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/middleware"
	utils2 "github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/gofiber/fiber/v2/utils"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

const (
	contextKey = "context"
)

func NewRestApp(
	ctx context.Context,
	cfg config.Config,
	route router.Router,
) *fiber.App {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Use absolute path for the templates directory
	templatePath := filepath.Join(wd, "/docs/swagger-ui")
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
			message := entity.Error
			errorMap := map[string]any{}

			// Retrieve the custom status code if it's a *fiber.Error
			var fiberError *fiber.Error
			slog.Info("Is fiber error", slog.Bool("is fiber error", errors.As(err, &fiberError)))
			if errors.As(err, &fiberError) {
				code = fiberError.Code
				status = utils.StatusMessage(fiberError.Code)

				if cfg.App.Env != "production" {
					message = fiberError.Error()
				}
			}

			// Retrieve the custom status code if it's an utils2.CustomError
			var customError *utils2.CustomError
			slog.Info("Is custom error: ", slog.Bool("is custom error", errors.As(err, &customError)))
			if errors.As(err, &customError) {
				code = customError.Code
				status = utils.StatusMessage(customError.Code)
				errorMap = customError.Fields

				if cfg.App.Env != "production" {
					message = customError.Error()
				}
			}

			return rest_util.ResponseGeneral(ctx, code, model.Response[any]{
				Code:    code,
				Status:  status,
				Message: message,
				Errors:  errorMap,
			})
		},
	})

	// Middleware to attach the context to the request
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals(contextKey, ctx)
		return c.Next()
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
	router.SetRoute(fiberApp, route)

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

func GetContext(ctx *fiber.Ctx) context.Context {
	return ctx.Locals(contextKey).(context.Context)
}
