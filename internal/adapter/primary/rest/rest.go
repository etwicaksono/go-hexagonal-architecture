package rest

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest/middleware"
	"github.com/etwicaksono/go-hexagonal-architecture/router"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"
)

const (
	ContextKey = "context"
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
	templatePath := filepath.Join(wd, "../docs/swagger-ui")
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

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				status = utils.StatusMessage(e.Code)

				if cfg.App.Env != "production" {
					message = e.Error()
				}
			}

			return ctx.Status(code).JSON(model.Response{
				Code:    code,
				Status:  status,
				Message: message,
			})
		},
	})

	// Middleware to attach the context to the request
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals(ContextKey, ctx)
		return c.Next()
	})

	if cfg.App.RestRecovery {
		fiberApp.Use(recover2.New(recover2.Config{
			EnableStackTrace: cfg.App.RestEnableStackTrace,
		})) // Panic Handler
	}

	// Middleware before route
	middleware.CorsMiddleware(fiberApp, cfg.App)

	// SetRoute
	router.SetRoute(fiberApp, route)

	// Static files
	docPath := filepath.Join(wd, "../docs")
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
	return ctx.Locals(ContextKey).(context.Context)
}
