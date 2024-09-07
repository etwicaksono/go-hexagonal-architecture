package rest

import (
	"github.com/gofiber/fiber/v2"
)

type ExampleHandlerInterface interface {
	SendTextMessage(ctx *fiber.Ctx) (err error)
	GetTextMessage(ctx *fiber.Ctx) (err error)
	SendMultimediaMessage(ctx *fiber.Ctx) (err error)
	GetMultimediaMessage(ctx *fiber.Ctx) (err error)
}
