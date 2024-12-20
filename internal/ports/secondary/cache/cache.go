package cache

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/model"
)

type AuthCacheInterface interface {
	SetAuthenticatedToken(ctx context.Context, accessKey string, cachedData model.AuthCachedData) (err error)
	GetAuthenticatedToken(ctx context.Context, accessKey string) (cachedData model.AuthCachedData, err error)
	DeleteAuthenticatedToken(ctx context.Context, accessKey string) (err error)
}
