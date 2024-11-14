package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a AuthenticationHandler) Register(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	payload := new(model.RegisterRequest)
	err = payload_util.BodyParser(ctx, payload)
	if err != nil {
		slog.ErrorContext(context, "Failed to parse RegisterRequest", slog.String(constants.Error, err.Error()))
		return
	}

	err = a.app.Register(context, payload.ToEntity())
	if err != nil {
		slog.ErrorContext(context, "Failed to register user", slog.String(constants.Error, err.Error()))
		return
	}

	return rest_util.ResponseOk(ctx, "Register user success")
}
