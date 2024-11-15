package infrastructure

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"time"
)

type adapterMongo struct {
	ctx           context.Context
	config        config.Config
	connectionURL string
	client        *mongo.Client
}

func NewMongoDb(
	ctx context.Context,
	config config.Config,
) infrastructure.DbInterface {
	return &adapterMongo{
		ctx:    ctx,
		config: config,
		connectionURL: fmt.Sprintf(
			"%s://%s:%s@%s/%s%s",
			config.Db.Protocol,
			config.Db.Username,
			config.Db.Password,
			config.Db.Address,
			config.Db.Name,
			config.Db.Option,
		),
	}
}

func (a *adapterMongo) Connect() error {
	clientOptions := options.Client().ApplyURI(a.connectionURL).
		SetMaxPoolSize(uint64(a.config.Db.MaxOpenConnections)).              // Max open connections
		SetMaxConnIdleTime(a.config.Db.MaxConnectionIdletime * time.Second). // Max connection idle time
		SetServerSelectionTimeout(10 * time.Second).                         // Timeout to find a server
		SetConnectTimeout(10 * time.Second).                                 // Timeout for initial connection
		SetSocketTimeout(30 * time.Second)                                   // Timeout for read/write on each socket

	client, err := mongo.Connect(a.ctx, clientOptions)
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to connect to MongoDB", slog.String("connection", a.connectionURL), slog.String(constants.Error, err.Error()))
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
		slog.ErrorContext(a.ctx, "Failed to disconnect to MongoDB", slog.String("connection", a.connectionURL), slog.String(constants.Error, err.Error()))
	}
}

func (a *adapterMongo) GetClient() *entity.DbClient {
	return &entity.DbClient{
		MongoClient: a.client,
	}
}
