package auth_cache

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
)

func (cache redisCache) DeleteAuthenticatedToken(ctx context.Context, accessKey string) (err error) {
	cacheKey := fmt.Sprintf("%s:%s", constants.AuthenticatedTokenRedisPrefix, accessKey)
	err = cache.Client.Del(ctx, cacheKey).Err()
	if err != nil {
		slog.ErrorContext(ctx, "failed to DeleteAuthenticatedToken", slog.String(constants.Error, err.Error()))
		return err
	}

	return nil
}
