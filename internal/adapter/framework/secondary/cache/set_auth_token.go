package cache

import (
	"context"
	"encoding/json"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
	"log/slog"
	"time"
)

func (cache redisCache) SetAuthToken(ctx context.Context, tokenKey string, token model.TokenData) (err error) {
	tokenByte, err := json.Marshal(token)
	if err != nil {
		slog.ErrorContext(ctx, "failed marshal auth token", slog.String("error", err.Error()))
		return
	}

	expiredAt := token.ExpiredDate.Sub(time.Now())

	err = cache.Client.Set(ctx, tokenKey, tokenByte, expiredAt).Err()
	if err != nil {
		slog.ErrorContext(ctx, "failed set auth token from redis", slog.String("error", err.Error()))
		return
	}

	return
}
