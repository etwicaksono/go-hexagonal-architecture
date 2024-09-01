package rest

import "github.com/gofiber/fiber/v2"

type SwaggerHandlerInterface interface {
	Swagger(ctx *fiber.Ctx) (err error)
}
