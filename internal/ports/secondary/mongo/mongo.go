package mongo

import "go.mongodb.org/mongo-driver/mongo"

type MongoInterface interface {
	Connect() error
	Disconnect()
	GetClient() *mongo.Client
}
