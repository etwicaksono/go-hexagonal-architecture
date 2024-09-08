package example_mongo

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	model2 "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func (e exampleMongo) InsertMultimediaMessage(ctx context.Context, objs []entity.MessageMultimediaItem) (entity.BulkWriteResult, error) {
	if len(objs) == 0 {
		return entity.BulkWriteResult{}, fmt.Errorf("no object to insert")
	}
	bulkCommands := make([]mongo.WriteModel, len(objs))
	collection := e.client.Database(e.dbName).Collection(e.collection)

	for i, obj := range objs {
		message := model2.FromMessageMultimediaItemEntity(obj)
		bulkCommands[i] = mongo.NewInsertOneModel().SetDocument(message)
	}

	result, err := collection.BulkWrite(ctx, bulkCommands)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to BulkWrite multimedia message", slog.String(entity.Error, err.Error()))
		return entity.BulkWriteResult{}, err
	}

	slog.InfoContext(
		ctx,
		"Successfully BulkWrite multimedia message",
		slog.Int("upserted count", int(result.UpsertedCount)),
		slog.Int("modified count", int(result.ModifiedCount)),
		slog.Int("objects count", len(objs)),
	)

	return entity.BulkWriteResult(*result), err
}
