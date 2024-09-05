package rest

import (
	"github.com/gofiber/fiber/v2"
)

type ExampleHandlerInterface interface {
	GetExample(ctx *fiber.Ctx) (err error)
}
