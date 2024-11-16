package example_message_core

import (
	"context"
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"log/slog"
)

func (e exampleMessageCore) GetMultimediaMessage(ctx context.Context) ([]entity.MessageMultimediaItem, error) {
	messages, err := e.db.FindAllMultimediaMessage(ctx)
	if err != nil && !errors.Is(err, errors2.ErrNoData) {
		slog.ErrorContext(ctx, "Failed to find all multimedia message", slog.String(constants.Error, err.Error()))
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

			switch file.Storage {
			case valueobject.MultimediaStorage_MINIO:
				{
					protocol := "http"
					if e.minio.IsUseSSL() {
						protocol = "https"
					}
					fileResult.File = fmt.Sprintf("%s://%s/%s/%s", protocol, e.minio.GetEndpoint(), e.minio.GetBucketName(), file.File)
				}
			case valueobject.MultimediaStorage_LOCAL:
				{
					fileResult.File = fmt.Sprintf("%s:%d/%s", e.appConfig.RestHost, e.appConfig.RestPort, file.File)
				}
			}

			msgResult.Files = append(msgResult.Files, fileResult)
		}
		result = append(result, msgResult)
	}
	return result, nil
}
