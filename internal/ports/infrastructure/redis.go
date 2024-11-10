package infrastructure

import "github.com/redis/go-redis/v9"

type RedisInterface interface {
	Connect()
	Disconnect()
	GetClient() (redisClient *redis.Client)
}
