package auth_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"
)

func (cache redisCache) GetAuthenticatedToken(ctx context.Context, accessKey string) (cachedData model.AuthCachedData, err error) {
	cacheKey := fmt.Sprintf("%s:%s", constants.AuthenticatedTokenRedisPrefix, accessKey)
	result, err := cache.Client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		slog.ErrorContext(ctx, "failed to GetAuthenticatedToken from redis", slog.String("cacheKey", cacheKey), slog.String(constants.Error, err.Error()))
		return
	}

	err = json.Unmarshal(result, &cachedData)
	if err != nil {
		slog.ErrorContext(ctx, "failed to unmarshal AuthCachedData", slog.String("cacheKey", cacheKey), slog.String(constants.Error, err.Error()))
		return
	}

	return
}
