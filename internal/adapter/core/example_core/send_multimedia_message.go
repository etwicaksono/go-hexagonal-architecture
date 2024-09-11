package example_core

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/string_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/validation_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (e exampleCore) SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) error {
	var files []entity.FileItem

	for _, requestFile := range request.Files {
		//Validate extension
		allowedTypes := []string{".jpg", ".jpeg", ".png", ".txt"}
		if !validation_util.IsValidExtension(allowedTypes, requestFile.Filename) {
			return error_util.ValidationError(
				fiber.Map{
					"files": fmt.Sprintf(
						"invalid file type (%s). Allowed types are %s",
						requestFile.Filename,
						string_util.Implode(allowedTypes, ", "),
					),
				},
			)
		}

		ext := filepath.Ext(requestFile.Filename)
		fileNameNoExtension := strings.TrimSuffix(requestFile.Filename, ext)
		fileName := fmt.Sprintf("%s-%d%s", payload_util.Slugify(fileNameNoExtension), time.Now().UnixNano(), ext)
		switch request.Storage {
		case entity.MultimediaStorage_LOCAL:
			{
				path := fmt.Sprintf("uploaded/%s", fileName)
				file, err := os.Create(path)
				if err != nil {
					return err
				}
				closeFile := func(file *os.File) {
					err := file.Close()
					if err != nil {
						slog.ErrorContext(ctx, "Failed to close file", slog.String("path", path), slog.String(entity.Error, err.Error()))
					}
				}

				// Write the file data
				_, err = file.Write(requestFile.Data)
				if err != nil {
					closeFile(file)
					return err
				}
				closeFile(file)
				files = append(
					files,
					entity.FileItem{File: path, Storage: entity.MultimediaStorage_name[int32(entity.MultimediaStorage_LOCAL)]},
				)
			}
		case entity.MultimediaStorage_MINIO:
			{
				filePath := fmt.Sprint(constants.MINIO_EXAMPLE_MESSAGE_PATH, "/", fileName)
				info, err := e.minio.Upload(ctx, requestFile.Data, requestFile.ContentType, filePath)
				if err != nil {
					slog.ErrorContext(ctx, "Failed to upload file", slog.String("path", filePath), slog.String(entity.Error, err.Error()))
					return err
				}

				files = append(
					files,
					entity.FileItem{File: info.Key, Storage: entity.MultimediaStorage_name[int32(entity.MultimediaStorage_MINIO)]},
				)
			}
		default:
			{
				return error_util.ValidationError(fiber.Map{"storage": "invalid storage type"})
			}
		}
	}

	objs := []entity.MessageMultimediaItem{
		{
			Sender:   request.Sender,
			Receiver: request.Receiver,
			Message:  request.Message,
			Files:    files,
		},
	}
	_, err := e.db.InsertMultimediaMessage(ctx, objs)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to insert multimedia message", slog.String(entity.Error, err.Error()))
		return err
	}

	return nil
}
