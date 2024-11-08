package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a AuthenticationHandler) Login(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	payload := new(model.LoginRequest)
	err = payload_util.BodyParser(ctx, payload)
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(context, "Failed to parse LoginRequest", slog.String(entity.Error, err.Error()))
		}
		return
	}

	tokenGenerated, err := a.app.Login(context, payload.ToEntity())
	if err != nil {
		if error_util.IsRealError(err) {
			slog.ErrorContext(context, "Failed to login user", slog.String(entity.Error, err.Error()))
		}
		return
	}

	return rest_util.ResponseOkWithData(ctx, model.FromTokenGeneratedEntity(tokenGenerated), "authenticated")
}
