package injector

import (
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func fiberProvider(cfg config.Config) *fiber.App {
	return fiber.New(fiber.Config{
		IdleTimeout:  cfg.App.IdleTimeout,
		WriteTimeout: cfg.App.WriteTimeout,
		ReadTimeout:  cfg.App.ReadTimeout,
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
}
