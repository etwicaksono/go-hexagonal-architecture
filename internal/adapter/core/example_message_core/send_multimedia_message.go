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
	var tempFiles []entity.FileItem
	var resultFiles []entity.FileItem
	deleteLocalTempFiles := func(tempFiles *[]entity.FileItem) error {
		for _, file := range *tempFiles {
			err := os.Remove(file.File)
			if err != nil {
				slog.ErrorContext(ctx, "Failed to delete file", slog.String("path", file.File), slog.String(entity.Error, err.Error()))
				return err
			}
		}
		return nil
	}
	defer func() {
		err := deleteLocalTempFiles(&tempFiles)
		if err != nil {
			slog.ErrorContext(ctx, errorsConst.ErrFailedToDeleteTempFiles.Error(), slog.String(entity.Error, err.Error()))
		}
	}()

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
				pathDir := "./uploaded/temp"
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
					slog.ErrorContext(ctx, "Failed to write file data", slog.String("path", pathFile), slog.String(entity.Error, err.Error()))
					return err
				}
				closeFile(file)
				tempFiles = append(
					tempFiles,
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

				tempFiles = append(
					tempFiles,
					entity.FileItem{File: info.Key, Storage: request.Storage.ToString()},
				)
			}
		default:
			{
				return error_util.ErrorValidation(fiber.Map{"storage": "invalid storage type"})
			}
		}
	}

	// TODO: move file from temp to correct directory
	switch request.Storage {
	case valueobject.MultimediaStorage_LOCAL:
		{
			for _, file := range tempFiles {
				currentFilePath := file.File
				newDir := "./uploaded/message"
				// Check if the new directory exists, if not, create it
				if _, err := os.Stat(newDir); os.IsNotExist(err) {
					err := os.MkdirAll(newDir, os.ModePerm)
					if err != nil {
						fmt.Println("Failed to create directory:", err)
						return err
					}
				}

				// Define the new file path in the target directory
				newFilePath := filepath.Join(newDir, filepath.Base(currentFilePath))

				// Move the file to the new directory
				if err := os.Rename(currentFilePath, newFilePath); err != nil {
					fmt.Println("Failed to move file:", err)
					return err
				}

				resultFiles = append(
					resultFiles,
					entity.FileItem{File: strings.Replace(newFilePath, "\\", "/", -1), Storage: request.Storage.ToString()},
				)
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
	_, err := e.db.InsertMultimediaMessage(ctx, objs)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to insert multimedia message", slog.String(entity.Error, err.Error()))
		return err
	}

	return nil
}
