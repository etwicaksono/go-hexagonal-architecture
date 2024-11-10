package middleware

import (
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
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
				return rest_util.ResponseGeneral(ctx, fiber.StatusUnprocessableEntity, model.Response[any]{
					Code:    fiber.StatusUnprocessableEntity,
					Status:  "error",
					Message: "Unprocessable Entity - Invalid input",
				})
			}

			return err
		}
		return nil
	})
}
