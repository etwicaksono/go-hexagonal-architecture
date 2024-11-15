package infrastructure

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type adapterRedis struct {
	ctx           context.Context
	client        *redis.Client
	connectionURL string
	config        config.RedisConfig
}

func NewRedis(ctx context.Context, config config.Config) infrastructure.RedisInterface { // TODO: should implement db interface
	return &adapterRedis{
		ctx:           ctx,
		connectionURL: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		config:        config.Redis,
	}
}

func (a *adapterRedis) Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     a.connectionURL,
		Username: a.config.Username,
		Password: a.config.Password,
		DB:       a.config.Db,
	})
	slog.InfoContext(a.ctx, "Redis connected", slog.String(
		"connected to", a.connectionURL,
	))
	a.client = client
}

func (a *adapterRedis) Disconnect() {
	err := a.client.Close()
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to disconnect to Redis", slog.String("connection", a.connectionURL), slog.String(constants.Error, err.Error()))
	}
}

func (a *adapterRedis) GetClient() (redisClient *redis.Client) {
	return a.client
}
