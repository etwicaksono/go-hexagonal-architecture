package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a adapter) GetMultimediaMessage(ctx *fiber.Ctx) (err error) {
	context := rest.GetContext(ctx)
	messages, err := a.app.GetMultimediaMessage(context)
	if err != nil {
		slog.ErrorContext(context, "Failed to get multimedia message", slog.String(entity.Error, err.Error()))
		return err
	}

	type messageItem struct {
		Id       string   `json:"id"`
		Sender   string   `json:"sender"`
		Receiver string   `json:"receiver"`
		Message  string   `json:"message"`
		FileUrls []string `json:"fileUrls"`
	}
	var data []messageItem
	for _, message := range messages {
		var fileUrls []string

		for _, file := range message.Files {
			fileUrls = append(fileUrls, file.File)
		}

		data = append(data, messageItem{
			Id:       message.Id,
			Sender:   message.Sender,
			Receiver: message.Receiver,
			Message:  message.Message,
			FileUrls: fileUrls,
		})
	}

	return rest_util.ResponseOkWithData(ctx, data, "Get multimedia message success")
}
