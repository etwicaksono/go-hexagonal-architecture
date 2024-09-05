package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/model"
	"github.com/gofiber/fiber/v2"
)

func (a adapter) GetExample(ctx *fiber.Ctx) (err error) {
	err = a.app.DoSomethingInApp()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Get example success",
	})
}
