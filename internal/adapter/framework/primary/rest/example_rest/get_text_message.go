package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a adapter) GetTextMessage(ctx *fiber.Ctx) (err error) {
	context := rest.GetContext(ctx)
	messages, err := a.app.GetTextMessage(context)
	if err != nil {
		slog.ErrorContext(context, "Failed to get text message", slog.String(entity.Error, err.Error()))
		return err
	}

	var modelMessages []model.MessageTextItem
	for _, message := range messages {
		modelMessages = append(modelMessages, model.FromMessageTextItemEntity(message))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response[[]model.MessageTextItem]{
		Code:    fiber.StatusOK,
		Status:  entity.Success,
		Message: "Get text message success",
		Data:    modelMessages,
	})
}
