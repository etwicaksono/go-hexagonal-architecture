package infrastructure

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type adapterRedis struct {
	ctx           context.Context
	client        *redis.Client
	connectionURL string
	config        RedisConfig
}

type RedisConfig struct {
	Db       int
	Host     string
	Port     int
	Username string
	Password string
}

func NewRedis(ctx context.Context, config RedisConfig) infrastructure.RedisInterface {
	return &adapterRedis{
		ctx:           ctx,
		connectionURL: fmt.Sprintf("%s:%d", config.Host, config.Port),
		config:        config,
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
		slog.ErrorContext(a.ctx, "Failed to disconnect to Redis", slog.String("connection", a.connectionURL), slog.String(entity.Error, err.Error()))
	}
}

func (a *adapterRedis) GetClient() (redisClient *redis.Client) {
	return a.client
}