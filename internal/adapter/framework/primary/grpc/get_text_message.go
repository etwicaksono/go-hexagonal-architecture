package grpc

import (
	"context"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) GetTextMessage(context.Context, *emptypb.Empty) (*example.GetTextMessageResponse, error) {
	messages, err := a.handler.ExampleApp.GetTextMessage()
	if err != nil {
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
