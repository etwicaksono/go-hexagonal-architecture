package example_mongo

import (
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type exampleMongo struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewExampleMongo(config config.Config, mongoClient *mongo.Client) db.ExampleDbInterface {
	return &exampleMongo{
		client:     mongoClient,
		dbName:     config.Db.Name,
		collection: config.Db.ExampleCollection,
	}
}
