package middleware

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		return rest_util.ResponseGeneral(ctx, fiber.StatusNotFound, model.Response[any]{
			Code:    fiber.ErrNotFound.Code,
			Status:  fiber.ErrNotFound.Message,
			Message: "Sorry, page not found!",
		})
	})
}
