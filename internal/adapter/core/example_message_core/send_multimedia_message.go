package example_message_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/storage_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (e exampleMessageCore) SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) error {
	var tempFiles []entity.FileItem
	var resultFiles []entity.FileItem
	defer func() {
		err := storage_util.DeleteTempFiles(ctx, &tempFiles)
		if err != nil {
			slog.ErrorContext(ctx, errorsConst.ErrFailedToDeleteTempFiles.Error(), slog.String(constants.Error, err.Error()))
		}
	}()

	err := validation_util.IsValidMultimediaFileExtension(request.Files, []string{".jpg", ".jpeg", ".png", ".txt"})
	if err != nil {
		return err
	}

	err = storage_util.SaveToTemp(ctx, request.Files, &tempFiles)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to save file to temp", slog.String(constants.Error, err.Error()))
		return err
	}

	switch request.Storage {
	case valueobject.MultimediaStorage_LOCAL:
		{
			resultFiles, err = storage_util.MoveFromTemp(storage_util.MoveFromTempArgs{
				Ctx:         ctx,
				TempFiles:   tempFiles,
				NewFilePath: "./uploaded/message",
				Storage:     valueobject.MultimediaStorage_LOCAL,
			})
			if err != nil {
				slog.ErrorContext(ctx, "Failed to move file from temp to local", slog.String(constants.Error, err.Error()))
				return err
			}
		}
	case valueobject.MultimediaStorage_MINIO:
		{
			resultFiles, err = storage_util.MoveFromTemp(storage_util.MoveFromTempArgs{
				Ctx:         ctx,
				TempFiles:   tempFiles,
				NewFilePath: constants.MinioExampleMessagePath,
				Storage:     valueobject.MultimediaStorage_MINIO,
				Minio:       e.minio,
			})
			if err != nil {
				slog.ErrorContext(ctx, "Failed to move file from temp to minio", slog.String(constants.Error, err.Error()))
				return err
			}
		}
	default:
		{
			return error_util.ErrorValidation(fiber.Map{"storage": "invalid storage type"})
		}
	}

	objs := []entity.MessageMultimediaItem{
		{
			Sender:   request.Sender,
			Receiver: request.Receiver,
			Message:  request.Message,
			Files:    resultFiles,
		},
	}
	_, err = e.db.InsertMultimediaMessage(ctx, objs)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to insert multimedia message", slog.String(constants.Error, err.Error()))
		return err
	}

	return nil
}
