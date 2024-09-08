package example_core

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/utils"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
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
		if !utils.IsValidExtension(allowedTypes, requestFile.Filename) {
			return error_util.ValidationError(
				fiber.Map{
					"files": fmt.Sprintf(
						"invalid file type (%s). Allowed types are %s",
						requestFile.Filename,
						utils.Implode(allowedTypes, ", "),
					),
				},
			)
		}
		switch request.Storage {
		case entity.MultimediaStorage_LOCAL:
			{
				// Create a new file with the original filename
				ext := filepath.Ext(requestFile.Filename)
				fileNameNoExtension := strings.TrimSuffix(requestFile.Filename, ext)
				path := fmt.Sprintf("uploaded/%s-%d%s", fileNameNoExtension, time.Now().UnixNano(), ext)
				file, err := os.Create(path)
				if err != nil {
					return err
				}

				// Write the file data
				_, err = file.Write(requestFile.Data)
				if err != nil {
					file.Close()
					return err
				}
				file.Close()
				files = append(
					files,
					entity.FileItem{File: path, Storage: entity.MultimediaStorage_name[int32(entity.MultimediaStorage_LOCAL)]},
				)
			}
		case entity.MultimediaStorage_MINIO:
			{
				files = append(
					files,
					entity.FileItem{File: "saved to minio", Storage: entity.MultimediaStorage_name[int32(entity.MultimediaStorage_MINIO)]},
				) // TODO
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
