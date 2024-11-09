package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"log/slog"
)

func (a authenticationApp) Refresh(ctx context.Context, accessKey string) (result entity.TokenGenerated, err error) {
	if accessKey == "" {
		return entity.TokenGenerated{}, errorsConst.ErrTokenInvalid
	}

	result, err = a.core.Refresh(ctx, accessKey)
	if err != nil {
		slog.ErrorContext(ctx, "Error on refreshing auth token", slog.String(entity.Error, err.Error()))
		return
	}

	return
}
