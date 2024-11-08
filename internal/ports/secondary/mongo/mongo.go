package mongo

import "go.mongodb.org/mongo-driver/mongo"

type MongoInterface interface {
	Connect() (err error)
	Disconnect()
	GetClient() (mongoClient *mongo.Client)
}
