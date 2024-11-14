package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a ExampleMessageHandler) GetMultimediaMessage(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()
	messages, err := a.app.GetMultimediaMessage(context)
	if err != nil {
		slog.ErrorContext(context, "Failed to get multimedia message", slog.String(constants.Error, err.Error()))
		return err
	}

	var data []model.MessageMultimediaItem
	for _, message := range messages {
		var fileUrls []string

		for _, file := range message.Files {
			fileUrls = append(fileUrls, file.File)
		}

		data = append(data, model.FromMessageMultimediaItemEntity(message))
	}

	return rest_util.ResponseOkWithData(ctx, data, "Get multimedia message success")
}
