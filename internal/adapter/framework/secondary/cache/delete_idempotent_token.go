package cache

import (
	"context"
	"log/slog"
)

func (cache redisCache) DeleteToken(ctx context.Context, tokenKey string) (err error) {
	err = cache.Client.Del(ctx, tokenKey).Err()
	if err != nil {
		slog.ErrorContext(ctx, "failed delete idempotent token", slog.String("err", err.Error()))
		return err
	}

	return nil
}
