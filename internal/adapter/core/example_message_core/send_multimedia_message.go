package example_message_core

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/string_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (e exampleMessageCore) SendMultimediaMessage(ctx context.Context, request entity.SendMultimediaMessageRequest) error {
	var files []entity.FileItem

	// TODO: save file to temporary directory, delete in the end of process, move to correct directory when process run correctly
	for _, requestFile := range request.Files {
		//Validate extension
		allowedTypes := []string{".jpg", ".jpeg", ".png", ".txt"}
		if !validation_util.IsValidExtension(allowedTypes, requestFile.Filename) {
			return error_util.ErrorValidation(
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
		case valueobject.MultimediaStorage_LOCAL:
			{
				pathDir := "uploaded/temp"
				// Check if the directory exists, if not, create it
				if _, err := os.Stat(pathDir); os.IsNotExist(err) {
					err := os.MkdirAll(pathDir, os.ModePerm)
					if err != nil {
						slog.ErrorContext(ctx, "Failed to create directory", slog.String("path", pathDir), slog.String(entity.Error, err.Error()))
						return errorsConst.ErrInternalServer
					}
				}

				pathFile := fmt.Sprintf("%s/%s", pathDir, fileName)

				file, err := os.Create(pathFile)
				if err != nil {
					return err
				}
				closeFile := func(file *os.File) {
					err := file.Close()
					if err != nil {
						slog.ErrorContext(ctx, "Failed to close file", slog.String("path", pathFile), slog.String(entity.Error, err.Error()))
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
					entity.FileItem{File: pathFile, Storage: request.Storage.ToString()},
				)
			}
		case valueobject.MultimediaStorage_MINIO:
			{
				filePath := fmt.Sprint(constants.MINIO_EXAMPLE_MESSAGE_PATH, "/", fileName)
				info, err := e.minio.Upload(ctx, requestFile.Data, requestFile.ContentType, filePath)
				if err != nil {
					slog.ErrorContext(ctx, "Failed to upload file", slog.String("path", filePath), slog.String(entity.Error, err.Error()))
					return err
				}

				files = append(
					files,
					entity.FileItem{File: info.Key, Storage: request.Storage.ToString()},
				)
			}
		default:
			{
				return error_util.ErrorValidation(fiber.Map{"storage": "invalid storage type"})
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
