package middleware

import (
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/gofiber/fiber/v2"
)

func UnprocessableEntityMiddleware(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err != nil {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code == fiber.StatusUnprocessableEntity {
				return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.Response[any]{
					Code:    fiber.StatusUnprocessableEntity,
					Status:  "error",
					Message: "Unprocessable Entity - Invalid input",
				})
			}

			// Default error handler for other status codes
			return ctx.Status(code).JSON(model.Response[any]{
				Code:    code,
				Status:  "error",
				Message: err.Error(),
			})
		}
		return nil
	})
}
