package auth_cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func (cache redisCache) IsBlacklistedToken(ctx context.Context, accessKey string) (isBlacklisted bool, err error) {
	cacheKey := fmt.Sprintf("%s:%s", constants.BlackListedTokenRedisPrefix, accessKey)
	_, err = cache.Client.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// is blacklisted
		isBlacklisted = true
	} else {
		if errors.Is(redis.Nil, err) {
			return false, nil
		}
		slog.ErrorContext(ctx, "failed to get blacklisted token from redis", slog.String("cacheKey", cacheKey), slog.String(constants.Error, err.Error()))
		return
	}
	return
}
