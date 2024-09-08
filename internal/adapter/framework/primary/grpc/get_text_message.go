package grpc

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"log/slog"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) GetTextMessage(ctx context.Context, _ *emptypb.Empty) (*example.GetTextMessageResponse, error) {
	messages, err := a.handler.ExampleApp.GetTextMessage(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get text message", slog.String(entity.Error, err.Error()))
		return nil, err
	}

	var data []*example.MessageTextItem
	for _, message := range messages {
		data = append(data, message.ToProto())
	}

	return &example.GetTextMessageResponse{
		Data: data,
	}, nil
}
