package example_mongo

import (
	"context"

	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/infrastructure"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type exampleMongo struct {
	client     *mongo.Client
	dbName     string
	collection string
}
type Config struct {
	Client     *mongo.Client
	DBName     string
	Collection string
}

func NewExampleMongo(config config.Config, mongo infrastructure.MongoInterface) db.ExampleDbInterface {
	return &exampleMongo{
		client:     mongo.GetClient(),
		dbName:     config.Db.Name,
		collection: config.Db.ExampleCollection,
	}
}

func (e exampleMongo) FindAllTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	//TODO implement me
	panic("implement me")
}

func (e exampleMongo) UpsertTextMessage(ctx context.Context, objs []entity.MessageTextItem) (entity.BulkWriteResult, error) {
	//TODO implement me
	panic("implement me")
}
