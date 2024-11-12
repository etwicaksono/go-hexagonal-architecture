package example_message_mongo

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type exampleMessageMongo struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewExampleMessageMongo(config config.Config, mongoClient *mongo.Client) db.ExampleMessageDbInterface {
	return &exampleMessageMongo{
		client:     mongoClient,
		dbName:     config.Db.Name,
		collection: config.Db.ExampleMessageCollection,
	}
}
