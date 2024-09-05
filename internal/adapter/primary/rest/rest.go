package rest

import (
	"context"
	"errors"
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
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  cfg.App.IdleTimeout,
		WriteTimeout: cfg.App.WriteTimeout,
		ReadTimeout:  cfg.App.ReadTimeout,
		Prefork:      cfg.App.RestPrefork,
		Views:        html.New("./docs/swagger-ui", ".gohtml"),
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
	staticFiles := map[string]string{
		"/docs": "docs",
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
