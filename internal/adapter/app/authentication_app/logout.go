package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
)

func (a authenticationApp) Logout(ctx context.Context, authToken entity.AuthToken) (err error) {
	err = a.core.Logout(ctx, authToken)
	if err != nil {
		slog.ErrorContext(ctx, "Error on logging out", slog.String(entity.Error, err.Error()))
		return
	}

	return
}
