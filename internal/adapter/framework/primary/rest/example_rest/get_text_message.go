package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/gofiber/fiber/v2"
)

func (a adapter) GetTextMessage(ctx *fiber.Ctx) (err error) {
	context := rest.GetContext(ctx)
	messages, err := a.app.GetTextMessage(context)
	if err != nil {
		return err
	}

	var modelMessages []model.MessageTextItem
	for _, message := range messages {
		modelMessages = append(modelMessages, model.FromEntity(message))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response[[]model.MessageTextItem]{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "Get text message success",
		Data:    modelMessages,
	})
}
