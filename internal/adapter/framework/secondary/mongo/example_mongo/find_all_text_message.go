package example_mongo

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (e exampleMongo) FindAllTextMessage(ctx context.Context) ([]entity.MessageTextItem, error) {
	collection := e.client.Database(e.dbName).Collection(e.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []entity.MessageTextItem
	for cursor.Next(ctx) {
		var message entity.MessageTextItem
		if err = cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
