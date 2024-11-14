package example_message_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
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
		slog.ErrorContext(context, "Failed to parse RegisterRequest", slog.String(constants.Error, err.Error()))
		return
	}

	err = a.app.SendTextMessage(context, payload.ToEntity())
	if err != nil {
		slog.ErrorContext(context, "Failed to send text message", slog.String(constants.Error, err.Error()))
		return err
	}

	return rest_util.ResponseOk(ctx, "Send text message success")
}
