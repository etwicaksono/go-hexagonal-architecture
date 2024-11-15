package grpc

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"log/slog"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) GetMultimediaMessage(ctx context.Context, _ *emptypb.Empty) (*example.GetMultimediaMessageResponse, error) {
	messages, err := a.handler.ExampleMessageApp.GetMultimediaMessage(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get multimedia message", slog.String(constants.Error, err.Error()))
		return nil, err
	}

	var data []*example.MessageMultimediaItem
	for _, message := range messages {
		data = append(data, message.ToProto())
	}

	return &example.GetMultimediaMessageResponse{
		Data: data,
	}, nil
}
