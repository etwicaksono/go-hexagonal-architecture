package example_user_mongo

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/repository/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type exampleUserMongo struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewExampleUserMongo(config config.Config, mongoClient *mongo.Client) db.ExampleUserDbInterface {
	return &exampleUserMongo{
		client:     mongoClient,
		dbName:     config.Db.Name,
		collection: config.Db.ExampleMessageCollection,
	}
}
