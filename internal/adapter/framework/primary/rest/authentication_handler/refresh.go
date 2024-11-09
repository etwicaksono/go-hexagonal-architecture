package authentication_handler

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/payload_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func (a AuthenticationHandler) Refresh(ctx *fiber.Ctx) (err error) {
	context := ctx.UserContext()

	payload := new(model.RefreshTokenRequest)
	err = payload_util.BodyParser(ctx, payload)
	if err != nil {
		slog.ErrorContext(context, "Failed to parse RefreshTokenRequest", slog.String(entity.Error, err.Error()))
		return
	}

	if payload.Token == "" {
		return errorsConst.ErrTokenInvalid
	}

	tokenReversed, err := a.jwt.ReverseJwtToken(payload.Token)
	if err != nil {
		return
	}

	if tokenReversed.TokenType != rest_util.RefreshTokenType {
		return errorsConst.ErrTokenInvalid
	}

	tokenGenerated, err := a.app.Refresh(context, tokenReversed.AccessKey)
	if err != nil {
		slog.ErrorContext(context, "Failed to refresh auth token", slog.String(entity.Error, err.Error()))
		return
	}

	return rest_util.ResponseOkWithData(ctx, model.FromTokenGeneratedEntity(tokenGenerated), "token refreshed")
}
