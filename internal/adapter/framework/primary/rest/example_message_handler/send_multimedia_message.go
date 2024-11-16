package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/valueobject"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a ExampleMessageHandler) SendMultimediaMessage(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	userData, err := a.jwt.GetAuthContextData(ctx)
	if err != nil {
		slog.ErrorContext(context, "Failed to get auth token", slog.String(constants.Error, err.Error()))
		return
	}

	payload := new(model.SendMultimediaMessageRequest)
	err = payload_util.BodyParser(ctx, payload)
	if err != nil {
		slog.ErrorContext(context, "Failed to parse RegisterRequest", slog.String(constants.Error, err.Error()))
		return
	}

	// Validate storage
	if err = valueobject.ValidateMultimediaStorageString(payload.Storage); err != nil {
		return
	}

	// Handle file upload
	parsedFiles, err := payload_util.MultipartFormParser(ctx, "files")
	if err != nil {
		slog.ErrorContext(context, "Failed to parse multipart form at 'files' field", slog.String(constants.Error, err.Error()))
		return
	}

	payload.Files = parsedFiles["files"]
	payload.Sender = userData.UserId

	err = a.app.SendMultimediaMessage(context, payload.ToEntity())
	if err != nil {
		slog.ErrorContext(context, "Failed to send multimedia message", slog.String(constants.Error, err.Error()))
		return err
	}

	return rest_util.ResponseOk(ctx, "Send multimedia message success")
}
