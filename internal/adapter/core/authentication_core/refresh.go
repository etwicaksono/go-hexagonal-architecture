package authentication_core

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"
)

func (a authenticationCore) Refresh(ctx context.Context, accessKey string) (result entity.TokenGenerated, err error) {
	generatedJwt, err := a.jwt.GenerateJwtToken(accessKey)
	if err != nil {
		slog.ErrorContext(ctx, "Error on generating jwt token", slog.String(entity.Error, err.Error()))
		return
	}

	// TODO: save access key to cache

	return generatedJwt.ToEntity(), nil
}
