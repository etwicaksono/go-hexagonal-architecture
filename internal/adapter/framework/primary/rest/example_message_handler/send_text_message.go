package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a ExampleMessageHandler) SendTextMessage(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	payload := new(model.SendTextMessageRequest)
	err = payload_util.BodyParser(ctx, payload)
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(context, "Failed to parse RegisterRequest", slog.String(entity.Error, err.Error()))
		}
		return
	}

	err = a.app.SendTextMessage(context, payload.ToEntity())
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(context, "Failed to send text message", slog.String(entity.Error, err.Error()))
		}
		return err
	}

	return rest_util.ResponseOk(ctx, "Send text message success")
}
