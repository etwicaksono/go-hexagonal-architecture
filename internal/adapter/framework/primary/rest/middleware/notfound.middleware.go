package middleware

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response[any]{
			Code:    fiber.ErrNotFound.Code,
			Status:  fiber.ErrNotFound.Message,
			Message: "Sorry, page not found!",
		})
	})
}
