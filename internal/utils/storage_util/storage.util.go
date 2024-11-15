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
	"sync"
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
	var wg sync.WaitGroup
	errChan := make(chan error, len(*tempFiles)) // buffered to handle potential errors from multiple goroutines

	for _, file := range *tempFiles {
		wg.Add(1)
		go func(f entity.FileItem) {
			defer wg.Done()
			if err := os.Remove(f.File); err != nil {
				slog.ErrorContext(ctx, "Failed to delete file", slog.String("path", f.File), slog.String(constants.Error, err.Error()))
				errChan <- err
			}
		}(file)
	}

	// Wait for all deletions to complete
	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	for err := range errChan {
		if err != nil {
			return err // return the first error encountered
		}
	}
	return nil
}

func SaveToTemp(ctx context.Context, multimediaFiles []entity.MultimediaFile, tempFiles *[]entity.FileItem) error {
	var wg sync.WaitGroup
	var mu sync.Mutex                                 // To safely append to tempFiles from multiple goroutines
	errChan := make(chan error, len(multimediaFiles)) // Buffered channel for error handling

	for _, requestFile := range multimediaFiles {
		wg.Add(1)
		go func(file entity.MultimediaFile) {
			defer wg.Done()

			ext := filepath.Ext(file.Filename)
			fileNameNoExtension := strings.TrimSuffix(file.Filename, ext)
			fileName := fmt.Sprintf("%s-%d%s", payload_util.Slugify(fileNameNoExtension), time.Now().UnixNano(), ext)
			pathDir := "./uploaded/temp"

			// Check if the directory exists, if not, create it
			if err := createDirectory(pathDir); err != nil {
				slog.ErrorContext(ctx, "Failed to create directory", slog.String("path", pathDir), slog.String(constants.Error, err.Error()))
				errChan <- errorsConst.ErrInternalServer
				return
			}

			pathFile := fmt.Sprintf("%s/%s", pathDir, fileName)
			fileHandle, err := os.Create(pathFile)
			if err != nil {
				errChan <- err
				return
			}
			defer func() {
				if closeErr := fileHandle.Close(); closeErr != nil {
					slog.ErrorContext(ctx, "Failed to close file", slog.String("path", pathFile), slog.String(constants.Error, closeErr.Error()))
				}
			}()

			// Write the file data
			if _, err = fileHandle.Write(file.Data); err != nil {
				slog.ErrorContext(ctx, "Failed to write file data", slog.String("path", pathFile), slog.String(constants.Error, err.Error()))
				errChan <- err
				return
			}

			// Safely append to tempFiles
			mu.Lock()
			*tempFiles = append(*tempFiles, entity.FileItem{File: pathFile, Storage: valueobject.MultimediaStorage_LOCAL.ToString()})
			mu.Unlock()
		}(requestFile)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Return any errors encountered
	for err := range errChan {
		if err != nil {
			return err
		}
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

	var wg sync.WaitGroup
	var mu sync.Mutex                                // To safely append to resultFiles from multiple goroutines
	errChan := make(chan error, len(args.TempFiles)) // Buffered channel for error handling

	for _, file := range args.TempFiles {
		wg.Add(1)
		go func(file entity.FileItem) {
			defer wg.Done()

			switch args.Storage {
			case valueobject.MultimediaStorage_LOCAL:
				currentFilePath := file.File
				// Ensure directory exists
				if err := createDirectory(args.NewFilePath); err != nil {
					slog.ErrorContext(args.Ctx, "Failed to create directory", slog.String("path", args.NewFilePath), slog.String(constants.Error, err.Error()))
					errChan <- err
					return
				}

				// Define the new file path
				newFilePath := filepath.Join(args.NewFilePath, filepath.Base(currentFilePath))

				// Move the file
				if err := os.Rename(currentFilePath, newFilePath); err != nil {
					slog.ErrorContext(args.Ctx, "Failed to move file", slog.String("path", currentFilePath), slog.String(constants.Error, err.Error()))
					errChan <- err
					return
				}

				// Append the result
				mu.Lock()
				resultFiles = append(resultFiles, entity.FileItem{File: strings.Replace(newFilePath, "\\", "/", -1), Storage: valueobject.MultimediaStorage_LOCAL.ToString()})
				mu.Unlock()

			case valueobject.MultimediaStorage_MINIO:
				if args.Minio == nil {
					errChan <- errorsConst.ErrMinioNotInitialized
					return
				}
				fileName := filepath.Base(file.File)
				filePath := fmt.Sprint(args.NewFilePath, "/", fileName)

				// Open the file
				tempFile, err := os.Open(file.File)
				if err != nil {
					slog.ErrorContext(args.Ctx, "Failed to open file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					errChan <- err
					return
				}
				defer func(tempFile *os.File) {
					err := tempFile.Close()
					if err != nil {
						slog.ErrorContext(args.Ctx, "Failed to close file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					}
				}(tempFile)

				// Read file data
				fileBytes, err := io.ReadAll(tempFile)
				if err != nil {
					slog.ErrorContext(args.Ctx, "Failed to read file", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					errChan <- err
					return
				}

				// Detect content type
				buffer := make([]byte, 512)
				if _, err := tempFile.ReadAt(buffer, 0); err != nil {
					slog.ErrorContext(args.Ctx, "Failed to read first bytes", slog.String("path", file.File), slog.String(constants.Error, err.Error()))
					errChan <- err
					return
				}
				contentType := http.DetectContentType(buffer)

				// Upload file
				info, err := args.Minio.Upload(args.Ctx, fileBytes, contentType, filePath)
				if err != nil {
					slog.ErrorContext(args.Ctx, "Failed to upload file", slog.String("path", filePath), slog.String(constants.Error, err.Error()))
					errChan <- err
					return
				}

				// Append the result
				mu.Lock()
				resultFiles = append(resultFiles, entity.FileItem{File: info.Key, Storage: valueobject.MultimediaStorage_MINIO.ToString()})
				mu.Unlock()

			default:
				errChan <- error_util.ErrorValidation(fiber.Map{"storage": "invalid storage type"})
				return
			}
		}(file)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}
	return resultFiles, nil
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
