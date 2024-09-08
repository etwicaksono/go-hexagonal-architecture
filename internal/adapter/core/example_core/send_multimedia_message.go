package example_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (e exampleCore) SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) error {
	var fileUrl string // TODO: get file url

	switch request.Storage {
	case entity.MultimediaStorage_LOCAL:
		{
			fileUrl = "saved to local"
		}
	case entity.MultimediaStorage_MINIO:
		{
			fileUrl = "saved to minio"
		}
	default:
		{
			return error_util.ValidationError(fiber.Map{"storage": "invalid storage type"})
		}

	}

	objs := []entity.MessageMultimediaItem{
		{
			Sender:   request.Sender,
			Receiver: request.Receiver,
			Message:  request.Message,
			FileUrl:  fileUrl,
		},
	}
	_, err := e.db.InsertMultimediaMessage(ctx, objs)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to insert multimedia message", slog.String(entity.Error, err.Error()))
		return err
	}

	return nil
}
