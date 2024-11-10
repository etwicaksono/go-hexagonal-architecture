package cache

import (
	"context"
	"encoding/json"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"log/slog"
)

func (cache redisCache) GetAuthToken(ctx context.Context, tokenKey string) (token model.TokenData, err error) {
	result, err := cache.Client.Get(ctx, tokenKey).Bytes()
	if err != nil {
		slog.ErrorContext(ctx, "failed get auth token from redis", slog.String("error", err.Error()))
		return
	}

	err = json.Unmarshal(result, &token)
	if err != nil {
		slog.ErrorContext(ctx, "failed unmarshal auth token", slog.String("err", err.Error()))
		return
	}

	return
}
