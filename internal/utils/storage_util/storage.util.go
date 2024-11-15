package storage_util

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/minio"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"github.com/gofiber/fiber/v2"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MoveFromTempArgs struct {
	Ctx         context.Context
	TempFiles   []entity.FileItem
	NewFilePath string
	Storage     valueobject.MultimediaStorage
	Minio       minio.MinioInterface
}

func DeleteTempFiles(ctx context.Context, tempFiles *[]entity.FileItem) error {
	for _, file := range *tempFiles {
		err := os.Remove(file.File)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to delete file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
			return err
		}
	}
	return nil
}

func SaveToTemp(ctx context.Context, multimediaFiles []entity.MultimediaFile, tempFiles *[]entity.FileItem) error {
	for _, requestFile := range multimediaFiles {
		ext := filepath.Ext(requestFile.Filename)
		fileNameNoExtension := strings.TrimSuffix(requestFile.Filename, ext)
		fileName := fmt.Sprintf("%s-%d%s", payload_util.Slugify(fileNameNoExtension), time.Now().UnixNano(), ext)
		pathDir := "./uploaded/temp"
		// Check if the directory exists, if not, create it
		err := createDirectory(pathDir)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to create directory", slog.String("path", pathDir), slog.String(constants.Error, err.Error()))
			return errorsConst.ErrInternalServer
		}

		pathFile := fmt.Sprintf("%s/%s", pathDir, fileName)

		file, err := os.Create(pathFile)
		if err != nil {
			return err
		}
		closeFile := func(file *os.File) {
			err := file.Close()
			if err != nil {
				slog.ErrorContext(ctx, "Failed to close file", slog.String("path", pathFile), slog.String(constants.Error, err.Error()))
			}
		}

		// Write the file data
		_, err = file.Write(requestFile.Data)
		if err != nil {
			closeFile(file)
			slog.ErrorContext(ctx, "Failed to write file data", slog.String("path", pathFile), slog.String(constants.Error, err.Error()))
			return err
		}
		closeFile(file)
		*tempFiles = append(
			*tempFiles,
			entity.FileItem{File: pathFile, Storage: valueobject.MultimediaStorage_LOCAL.ToString()},
		)
	}
	return nil
}

func MoveFromTemp(args MoveFromTempArgs) (resultFiles []entity.FileItem, err error) {
	if args.Ctx == nil {
		return nil, errorsConst.ErrInvalidContext
	}
	if len(args.TempFiles) == 0 {
		return nil, errorsConst.ErrInvalidTempFiles
	}
	if args.NewFilePath == "" {
		return nil, errorsConst.ErrInvalidNewFilePath
	}

	for _, file := range args.TempFiles {
		switch args.Storage {
		case valueobject.MultimediaStorage_LOCAL:
			{
				currentFilePath := file.File
				// Check if the directory exists, if not, create it
				err = createDirectory(args.NewFilePath)
				if err != nil {
					slog.ErrorContext(args.Ctx, "Failed to create directory", slog.String("path", args.NewFilePath), slog.String(constants.Error, err.Error()))
					return nil, err
				}

				// Define the new file path in the target directory
				newFilePath := filepath.Join(args.NewFilePath, filepath.Base(currentFilePath))

				// Move the file to the new directory
				if err = os.Rename(currentFilePath, newFilePath); err != nil {
					slog.ErrorContext(args.Ctx, "Failed to move file", slog.String("path", currentFilePath), slog.String(constants.Error, err.Error()))
					return nil, err
				}

				resultFiles = append(
					resultFiles,
					entity.FileItem{File: strings.Replace(newFilePath, "\\", "/", -1), Storage: valueobject.MultimediaStorage_LOCAL.ToString()},
				)
			}
		case valueobject.MultimediaStorage_MINIO:
			{
				if args.Minio == nil {
					return nil, errorsConst.ErrMinioNotInitialized
				}
				fileName := filepath.Base(file.File)
				filePath := fmt.Sprint(args.NewFilePath, "/", fileName)

				// Open the file
				tempFile, err := os.Open(file.File)
				if err != nil {
					slog.ErrorContext(args.Ctx, "Failed to open file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					return nil, err
				}
				closeTempFile := func() {
					err := tempFile.Close()
					if err != nil {
						return
					}
				}

				// Read the file into a []byte
				fileBytes, err := io.ReadAll(tempFile)
				if err != nil {
					closeTempFile()
					slog.ErrorContext(args.Ctx, "Failed to read all file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					return nil, err
				}

				// Reset the file pointer back to the start for further reading if needed
				_, err = tempFile.Seek(0, 0)
				if err != nil {
					closeTempFile()
					slog.ErrorContext(args.Ctx, "Failed to reset file pointer", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					return nil, err
				}

				// Read the first 512 bytes of the file to detect content type
				buffer := make([]byte, 512)
				_, err = tempFile.Read(buffer)
				if err != nil {
					closeTempFile()
					slog.ErrorContext(args.Ctx, "Failed to read first 512 bytes of the file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					return nil, err
				}

				// Detect the content type
				contentType := http.DetectContentType(buffer)

				info, err := args.Minio.Upload(args.Ctx, fileBytes, contentType, filePath)
				if err != nil {
					closeTempFile()
					slog.ErrorContext(args.Ctx, "Failed to upload file", slog.String("path", filePath), slog.String(constants.Error, err.Error()))
					return nil, err
				}
				closeTempFile()

				resultFiles = append(
					resultFiles,
					entity.FileItem{File: info.Key, Storage: valueobject.MultimediaStorage_MINIO.ToString()},
				)
			}
		default:
			{
				return nil, error_util.ErrorValidation(fiber.Map{"storage": "invalid storage type"})
			}

		}
	}
	return
}

func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
