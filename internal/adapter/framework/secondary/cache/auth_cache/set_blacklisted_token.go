package auth_cache

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
	"time"
)

func (cache redisCache) SetBlacklistedToken(ctx context.Context, accessKey string, expiredDate time.Time) (err error) {
	expiredAt := expiredDate.Sub(time.Now())
	cacheKey := fmt.Sprintf("%s:%s", constants.AuthenticatedTokenRedisPrefix, accessKey)
	err = cache.Client.Set(ctx, cacheKey, true, expiredAt).Err()
	if err != nil {
		slog.ErrorContext(ctx, "failed to SetBlacklistedToken to redis", slog.String(constants.Error, err.Error()))
		return
	}

	return
}
