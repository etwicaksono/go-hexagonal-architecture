package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a ExampleMessageHandler) GetTextMessage(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()
	messages, err := a.app.GetTextMessage(context)
	if err != nil {
		slog.ErrorContext(context, "Failed to get text message", slog.String(constants.Error, err.Error()))
		return err
	}

	var modelMessages []model.MessageTextItem
	for _, message := range messages {
		modelMessages = append(modelMessages, model.FromMessageTextItemEntity(message))
	}

	return rest_util.ResponseOkWithData(ctx, modelMessages, "Get text message success")
}
