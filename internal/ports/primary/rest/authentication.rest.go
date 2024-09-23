package rest

import "github.com/gofiber/fiber/v2"

type AuthenticationHandlerInterface interface {
	Register(ctx *fiber.Ctx) (err error)
	Login(ctx *fiber.Ctx) (err error)
	Logout(ctx *fiber.Ctx) (err error)
	Refresh(ctx *fiber.Ctx) (err error)
}
