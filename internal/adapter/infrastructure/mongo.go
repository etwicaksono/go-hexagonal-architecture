package infrastructure

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"time"
)

type adapterMongo struct {
	ctx           context.Context
	connectionURL string
	Client        *mongo.Client
	config        mongoConfig
}

type mongoConfig struct {
	Protocol        string
	Address         string
	Name            string
	Username        string
	Password        string
	MaxConnOpen     int
	MaxConnIdle     int
	MaxConnLifetime time.Duration
	Option          string
}

func NewMongo(
	ctx context.Context,
	config config.Config,
) infrastructure.MongoInterface {
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
			Protocol:        config.Db.Protocol,
			Address:         config.Db.Address,
			Name:            config.Db.Name,
			Username:        config.Db.Username,
			Password:        config.Db.Password,
			MaxConnOpen:     config.Db.MaxConnOpen,
			MaxConnIdle:     config.Db.MaxConnIdle,
			MaxConnLifetime: config.Db.MaxConnLifetime,
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
	a.Client = client

	return nil
}

func (a *adapterMongo) Disconnect() {
	err := a.Client.Disconnect(a.ctx)
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to disconnect to MongoDB", slog.String("connection", a.connectionURL), slog.String(entity.Error, err.Error()))
	}
}

func (a *adapterMongo) GetClient() *mongo.Client {
	return a.Client
}
