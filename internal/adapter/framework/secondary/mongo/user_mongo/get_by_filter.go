package user_mongo

import (
	"context"
	"errors"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/model"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	errors2 "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func (e userMongo) GetByFilter(ctx context.Context, filter entity.UserGetFilter) ([]entity.User, error) {
	collection := e.client.Database(e.dbName).Collection(e.collection)
	var pipeline bson.D

	if len(filter.IDs) > 0 {
		var ids []primitive.ObjectID
		for _, id := range filter.IDs {
			_id, _ := primitive.ObjectIDFromHex(id)
			ids = append(ids, _id)
		}
		pipeline = append(pipeline, bson.E{Key: "_id", Value: bson.M{"$in": ids}})
	}

	if len(filter.Emails) > 0 {
		var emailRegexes []interface{}
		for _, t := range filter.Emails {
			// Use case-insensitive regex
			emailRegexes = append(emailRegexes, primitive.Regex{Pattern: "^" + t + "$", Options: "i"})
		}
		pipeline = append(pipeline, bson.E{Key: "email", Value: bson.M{"$in": emailRegexes}})
	}

	if len(filter.Names) > 0 {
		var nameRegexes []interface{}
		for _, t := range filter.Names {
			// Use case-insensitive regex for partial matches
			nameRegexes = append(nameRegexes, primitive.Regex{Pattern: ".*" + t + ".*", Options: "i"})
		}
		pipeline = append(pipeline, bson.E{Key: "name", Value: bson.M{"$in": nameRegexes}})
	}

	if len(filter.Usernames) > 0 {
		var usernameRegexes []interface{}
		for _, t := range filter.Usernames {
			// Use case-insensitive regex
			usernameRegexes = append(usernameRegexes, primitive.Regex{Pattern: "^" + t + "$", Options: "i"})
		}
		pipeline = append(pipeline, bson.E{Key: "username", Value: bson.M{"$in": usernameRegexes}})
	}

	if filter.Active.Valid {
		pipeline = append(pipeline, bson.E{Key: "active", Value: filter.Active.Bool})
	}

	findOptions := options.Find()
	cursor, err := collection.Find(ctx, pipeline, findOptions)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors2.ErrNoData
		}
		slog.ErrorContext(ctx, "Failed to get user", slog.String(constants.Error, err.Error()))
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []entity.User
	for cursor.Next(ctx) {
		var user model.User
		if err = cursor.Decode(&user); err != nil {
			slog.ErrorContext(ctx, "Failed to decode message", slog.String(constants.Error, err.Error()))
			return nil, err
		}
		users = append(users, user.ToEntity())
	}
	return users, nil
}
