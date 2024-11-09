package infrastructure

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	mongo2 "github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"time"
)

type adapterMongo struct {
	ctx           context.Context
	connectionURL string
	client        *mongo.Client
	config        mongoConfig
}

type mongoConfig struct {
	protocol        string
	address         string
	name            string
	username        string
	password        string
	maxConnOpen     int
	maxConnIdle     int
	maxConnLifetime time.Duration
	option          string
}

func NewMongo(
	ctx context.Context,
	config config.Config,
) mongo2.MongoInterface {
	return &adapterMongo{
		ctx: ctx,
		connectionURL: fmt.Sprintf(
			"%s://%s:%s@%s/%s%s",
			config.Db.Protocol,
			config.Db.Username,
			config.Db.Password,
			config.Db.Address,
			config.Db.Name,
			config.Db.Option,
		),
		config: mongoConfig{
			protocol:        config.Db.Protocol,
			address:         config.Db.Address,
			name:            config.Db.Name,
			username:        config.Db.Username,
			password:        config.Db.Password,
			maxConnOpen:     config.Db.MaxConnOpen,
			maxConnIdle:     config.Db.MaxConnIdle,
			maxConnLifetime: config.Db.MaxConnLifetime,
		},
	}
}

func (a *adapterMongo) Connect() error {
	clientOptions := options.Client().
		ApplyURI(a.connectionURL)

	client, err := mongo.Connect(a.ctx, clientOptions)
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to connect to MongoDB", slog.String("connection", a.connectionURL), slog.String(entity.Error, err.Error()))
		return err
	}

	slog.InfoContext(a.ctx, "MongoDB connected", slog.String(
		"connected to", a.connectionURL,
	))
	a.client = client

	return nil
}

func (a *adapterMongo) Disconnect() {
	err := a.client.Disconnect(a.ctx)
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to disconnect to MongoDB", slog.String("connection", a.connectionURL), slog.String(entity.Error, err.Error()))
	}
}

func (a *adapterMongo) GetClient() *mongo.Client {
	return a.client
}
