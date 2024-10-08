package example_rest

import (
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/string_util"
	"github.com/gofiber/fiber/v2"
	"io"
	"log/slog"
)

func (a adapter) SendMultimediaMessage(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	payload := new(model.SendMultimediaMessageRequest)
	err = ctx.BodyParser(payload)
	if err != nil {
		errParsing, errOther := payload_util.HandleParsingError(err)
		if errOther != nil {
			slog.ErrorContext(context, errOther.Error())
			return errOther
		}
		return error_util.ValidationError(errParsing)
	}

	// Validate storage
	switch payload.Storage {
	case model.MultimediaStorage_LOCAL, model.MultimediaStorage_MINIO:
		{
			// do nothing
		}
	default:
		{
			return error_util.ValidationError(
				fiber.Map{"storage": fmt.Sprintf(
					"Invalid storage type. Available types are: %s",
					string_util.Implode(
						[]string{
							string(model.MultimediaStorage_LOCAL),
							string(model.MultimediaStorage_MINIO)},
						", ",
					))},
			)
		}
	}

	// Handle file upload
	// Parse the multipart form containing the files
	form, err := ctx.MultipartForm()
	if err != nil {
		return error_util.ValidationError(fiber.Map{"files": "Failed to parse multipart form"})
	}

	// Get the files from the "files" field in the form
	files := form.File["files"]
	if files != nil {
		for _, file := range files {
			// Open the file
			openedFile, err := file.Open()
			if err != nil {
				return error_util.ValidationError(fiber.Map{"files": fmt.Sprintf("Failed to open file (%s)", file.Filename)})
			}
			defer func() {
				if err := openedFile.Close(); err != nil {
					slog.ErrorContext(context, fmt.Sprintf("Failed to close file (%s)", file.Filename), slog.String(entity.Error, err.Error()))
				}
			}()

			// Read file content into []byte
			fileBytes, err := io.ReadAll(openedFile)
			if err != nil {
				return error_util.ValidationError(fiber.Map{"files": fmt.Sprintf("Failed to read file (%s)", file.Filename)})
			}

			payload.Files = append(payload.Files, entity.MultimediaFile{
				Filename:    file.Filename,
				ContentType: file.Header.Get("Content-Type"),
				Data:        fileBytes,
			})
		}
	}

	err = a.app.SendMultimediaMessage(context, payload.ToEntity())
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(context, "Failed to send multimedia message", slog.String(entity.Error, err.Error()))
		}
		return err
	}

	return rest_util.ResponseOk(ctx, "Send multimedia message success")
}
