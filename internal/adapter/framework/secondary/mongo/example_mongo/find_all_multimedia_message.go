package example_mongo

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	model2 "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func (e exampleMongo) FindAllMultimediaMessage(ctx context.Context) ([]entity.MessageMultimediaItem, error) {
	collection := e.client.Database(e.dbName).Collection(e.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrNoData
		}
		slog.ErrorContext(ctx, "Failed to find all multimedia message", slog.String(entity.Error, err.Error()))
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []entity.MessageMultimediaItem
	for cursor.Next(ctx) {
		var message model2.MessageMultimediaItem
		if err = cursor.Decode(&message); err != nil {
			slog.ErrorContext(ctx, "Failed to decode message", slog.String(entity.Error, err.Error()))
			return nil, err
		}
		messages = append(messages, message.ToEntity())
	}
	return messages, nil
}
