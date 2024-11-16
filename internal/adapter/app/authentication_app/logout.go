package authentication_app

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
)

func (a authenticationApp) Logout(ctx context.Context, accessKey string) (err error) {
	err = a.core.Logout(ctx, accessKey)
	if err != nil {
		slog.ErrorContext(ctx, "Error on logging out", slog.String(constants.Error, err.Error()))
		return
	}

	return
}
