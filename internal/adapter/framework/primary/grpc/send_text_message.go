package grpc

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) SendTextMessage(ctx context.Context, request *example.SendTextMessageRequest) (*emptypb.Empty, error) {
	err := a.handler.ExampleApp.SendTextMessage(ctx, entity.SendTextMessageRequest{
		Sender:   request.Sender,
		Receiver: request.Receiver,
		Message:  request.Message,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
