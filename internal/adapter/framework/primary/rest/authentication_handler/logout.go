package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a AuthenticationHandler) Logout(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()
	authToken, err := a.jwt.GetJwtAuthToken(ctx)
	if err != nil {
		slog.ErrorContext(context, "Failed to get auth token", slog.String(entity.Error, err.Error()))
		return
	}

	err = a.app.Logout(context, *authToken)
	if err != nil {
		slog.ErrorContext(context, "Failed to logout", slog.String(entity.Error, err.Error()))
		return
	}

	return rest_util.ResponseOk(ctx, "Logout success")
}
