package rest_util

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/gofiber/fiber/v2"
)

func ResponseGeneral[T any](
	ctx *fiber.Ctx,
	code int,
	response model.Response[T],
) error {
	return ctx.Status(code).JSON(response)
}

func ResponseOk(
	ctx *fiber.Ctx,
	message string,
) error {
	return ctx.Status(fiber.StatusOK).JSON(model.Response[any]{
		Code:    fiber.StatusOK,
		Status:  entity.Success,
		Message: message,
	})
}

func ResponseOkWithData[T any](
	ctx *fiber.Ctx,
	data T,
	message string,
) error {
	return ctx.Status(fiber.StatusOK).JSON(model.Response[T]{
		Code:    fiber.StatusOK,
		Status:  entity.Success,
		Message: message,
		Data:    data,
	})
}
