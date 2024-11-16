package auth_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
	"time"
)

func (cache redisCache) SetAuthenticatedToken(ctx context.Context, accessKey string, cachedData model.AuthCachedData) (err error) {
	tokenByte, err := json.Marshal(cachedData)
	if err != nil {
		slog.ErrorContext(ctx, "failed marshal authenticated token", slog.String(constants.Error, err.Error()))
		return
	}

	expiredAt := cachedData.ExpiredAt.Sub(time.Now())

	cacheKey := fmt.Sprintf("%s:%s", constants.AuthenticatedTokenRedisPrefix, accessKey)
	err = cache.Client.Set(ctx, cacheKey, tokenByte, expiredAt).Err()
	if err != nil {
		slog.ErrorContext(ctx, "failed set authenticated token to redis", slog.String(constants.Error, err.Error()))
		return
	}

	return
}
