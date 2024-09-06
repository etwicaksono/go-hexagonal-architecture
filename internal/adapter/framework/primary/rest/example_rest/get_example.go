package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/gofiber/fiber/v2"
)

func (a adapter) GetExample(ctx *fiber.Ctx) (err error) {
	messages, err := a.app.GetTextMessage()
	if err != nil {
		return err
	}

	var modelMessages []model.MessageTextItem
	for _, message := range messages {
		modelMessages = append(modelMessages, *message.ToModel())
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response[[]model.MessageTextItem]{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Get example success",
		Data:    modelMessages,
	})
}
