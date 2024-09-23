package example_message_core

import (
	"context"
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"log/slog"
)

func (e exampleMessageCore) GetMultimediaMessage(ctx context.Context) ([]entity.MessageMultimediaItem, error) {
	messages, err := e.db.FindAllMultimediaMessage(ctx)
	if err != nil && !errors.Is(err, errors2.ErrNoData) {
		slog.ErrorContext(ctx, "Failed to find all multimedia message", slog.String(entity.Error, err.Error()))
		return nil, err
	}

	var result []entity.MessageMultimediaItem
	for _, message := range messages {
		msgResult := entity.MessageMultimediaItem{
			Id:       message.Id,
			Sender:   message.Sender,
			Receiver: message.Receiver,
			Message:  message.Message,
		}
		for _, file := range message.Files {
			fileResult := file
			protocol := "http"
			if e.minio.IsUseSSL() {
				protocol = "https"
			}
			fileResult.File = fmt.Sprintf("%s://%s/%s/%s", protocol, e.minio.GetEndpoint(), e.minio.GetBucketName(), file.File)
			msgResult.Files = append(msgResult.Files, fileResult)
		}
		result = append(result, msgResult)
	}
	return result, nil
}