package user_mongo

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/secondary/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewUserMongo(config config.Config, mongoClient *mongo.Client) db.UserDbInterface {
	return &userMongo{
		client:     mongoClient,
		dbName:     config.Db.Name,
		collection: config.Db.ExampleUserCollection,
	}
}
