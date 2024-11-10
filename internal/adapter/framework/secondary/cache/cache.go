package cache

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/cache"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	*redis.Client
}

func NewCache(redisClient *redis.Client) cache.CacheInterface {
	return &redisCache{
		redisClient,
	}
}
