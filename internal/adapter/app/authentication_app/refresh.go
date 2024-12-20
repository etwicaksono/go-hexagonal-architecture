package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"log/slog"
)

func (a authenticationApp) Refresh(ctx context.Context, request entity.RefreshTokenRequest) (result entity.TokenGenerated, err error) {
	err = validation_util.ValidateStruct(a.validator, model.FromAuthTokenRequestEntity(request))
	if err != nil {
		return
	}

	tokenReversed, err := a.jwt.ReverseJwtToken(request.Token)
	if err != nil {
		return
	}

	if tokenReversed.TokenType != rest_util.RefreshTokenType {
		return result, errorsConst.ErrTokenInvalid
	}

	result, err = a.core.Refresh(ctx, tokenReversed.AccessKey)
	if err != nil {
		slog.ErrorContext(ctx, "Error on refreshing auth token", slog.String(constants.Error, err.Error()))
		return
	}

	return
}
