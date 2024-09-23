package example_message_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a adapter) GetTextMessage(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()
	messages, err := a.app.GetTextMessage(context)
	if err != nil {
		slog.ErrorContext(context, "Failed to get text message", slog.String(entity.Error, err.Error()))
		return err
	}

	var modelMessages []model.MessageTextItem
	for _, message := range messages {
		modelMessages = append(modelMessages, model.FromMessageTextItemEntity(message))
	}

	return rest_util.ResponseOkWithData(ctx, modelMessages, "Get text message success")
}
