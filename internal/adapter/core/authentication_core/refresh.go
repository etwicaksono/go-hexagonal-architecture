package authentication_core

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func (a authenticationCore) Refresh(ctx context.Context, accessKey string) (result entity.TokenGenerated, err error) {
	// Get cached auth data
	userData, err := a.cache.GetAuthenticatedToken(ctx, accessKey)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return result, errorsConst.ErrTokenInvalid
		}
		return entity.TokenGenerated{}, err
	}
	// Generate new access key
	newAccessKey, err := utils.PasswordGenerate(userData.UserId)
	if err != nil {
		slog.ErrorContext(ctx, "Error on generating access key", slog.String(constants.Error, err.Error()))
		return
	}

	generatedJwt, err := a.jwt.GenerateJwtToken(newAccessKey)
	if err != nil {
		slog.ErrorContext(ctx, "Error on generating jwt token", slog.String(constants.Error, err.Error()))
		return
	}

	// Delete old token
	err = a.cache.DeleteAuthenticatedToken(ctx, accessKey)
	if err != nil {
		return
	}

	// Save new token
	err = a.cache.SetAuthenticatedToken(ctx, newAccessKey, model.AuthCachedData{
		UserId:    userData.UserId,
		AccessKey: newAccessKey,
		ExpiredAt: generatedJwt.ExpiredAt,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Error on saving auth token to cache", slog.String(constants.Error, err.Error()))
		return
	}

	return generatedJwt.ToEntity(), nil
}
