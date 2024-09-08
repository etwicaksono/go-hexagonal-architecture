package example_rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/utils"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"io"
	"log/slog"
)

func (a adapter) SendMultimediaMessage(ctx *fiber.Ctx) (err error) {
	context := rest.GetContext(ctx)

	payload := new(model.SendMultimediaMessageRequest)
	err = ctx.BodyParser(payload)
	if err != nil {
		errParsing, errOther := utils.HandleParsingError(err)
		if errOther != nil {
			slog.ErrorContext(context, errOther.Error())
			return errOther
		}
		return error_util.ValidationError(errParsing)
	}

	// Handle file upload
	fileHeader, err := ctx.FormFile("file")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		// Read the file content into a byte slice
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return error_util.ValidationError(fiber.Map{"file": "Failed to read file"})
		}

		payload.Files = append(payload.Files, entity.MultimediaFile{
			Filename: fileHeader.Filename,
			Data:     fileBytes,
		})
	}

	err = a.app.SendMultimediaMessage(context, payload.ToEntity())
	if err != nil {
		if !error_util.IsValidationError(err) {
			slog.ErrorContext(context, "Failed to send multimedia message", slog.String(entity.Error, err.Error()))
		}
		return err
	}

	return rest_util.ResponseOk(ctx, "Send multimedia message success")
}
