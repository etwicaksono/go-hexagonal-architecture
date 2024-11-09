package cache

import (
	"context"
	"log/slog"
	"time"
)

func (cache redisCache) SetIdempotentToken(ctx context.Context, tokenKey string) (bool, error) {
	success, err := cache.Client.SetNX(ctx, tokenKey, nil, 5*time.Minute).Result()
	if err != nil {
		slog.ErrorContext(ctx, "failed set idempotent token", slog.String("error", err.Error()), err)
		return true, err
	}

	return !success, nil
}
