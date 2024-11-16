package auth_cache

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/cache"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	*redis.Client
}

func NewCache(redisClient *redis.Client) cache.AuthCacheInterface {
	return &redisCache{
		redisClient,
	}
}
