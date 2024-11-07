package user_mongo

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/model"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func (e userMongo) GetByFilter(ctx context.Context, filter entity.UserGetFilter) ([]entity.User, error) {
	collection := e.client.Database(e.dbName).Collection(e.collection)
	var pipeline []bson.D

	if len(filter.IDs) > 0 {
		var ids []primitive.ObjectID
		for _, id := range filter.IDs {
			_id, _ := primitive.ObjectIDFromHex(id)
			ids = append(ids, _id)
		}
		pipeline = append(pipeline, bson.D{{"_id", bson.M{"$in": ids}}})
	}

	if len(filter.Emails) > 0 {
		var emailRegexes []interface{}
		for _, t := range filter.Emails {
			// Use case-insensitive regex
			emailRegexes = append(emailRegexes, primitive.Regex{Pattern: "^" + t + "$", Options: "i"})
		}
		pipeline = append(pipeline, bson.D{{"email", bson.M{"$in": emailRegexes}}})
	}

	if len(filter.Names) > 0 {
		var nameRegexes []interface{}
		for _, t := range filter.Names {
			// Use case-insensitive regex for partial matches
			nameRegexes = append(nameRegexes, primitive.Regex{Pattern: ".*" + t + ".*", Options: "i"})
		}
		pipeline = append(pipeline, bson.D{{"name", bson.M{"$in": nameRegexes}}})
	}

	if len(filter.Usernames) > 0 {
		var usernameRegexes []interface{}
		for _, t := range filter.Usernames {
			// Use case-insensitive regex
			usernameRegexes = append(usernameRegexes, primitive.Regex{Pattern: "^" + t + "$", Options: "i"})
		}
		pipeline = append(pipeline, bson.D{{"username", bson.M{"$in": usernameRegexes}}})
	}

	if filter.Active != nil {
		pipeline = append(pipeline, bson.D{{"active", *filter.Active}})
	}

	cursor, err := collection.Find(ctx, pipeline)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors2.ErrNoData
		}
		slog.ErrorContext(ctx, "Failed to get user", slog.String(entity.Error, err.Error()))
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []entity.User
	for cursor.Next(ctx) {
		var message model.User
		if err = cursor.Decode(&message); err != nil {
			slog.ErrorContext(ctx, "Failed to decode message", slog.String(entity.Error, err.Error()))
			return nil, err
		}
		users = append(users, message.ToEntity())
	}
	return users, nil
}