package cache

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
)

type CacheInterface interface {
	GetAuthToken(ctx context.Context, tokenKey string) (token model.TokenData, err error)
	SetToken(ctx context.Context, tokenKey string, token model.TokenData) (err error)
	SetIdempotentToken(ctx context.Context, tokenKey string) (success bool, err error)
	DeleteToken(ctx context.Context, tokenKey string) (err error)
}
