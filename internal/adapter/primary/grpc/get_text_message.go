package grpc

import (
	"context"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) GetTextMessage(context.Context, *emptypb.Empty) (*example.GetTextMessageResponse, error) {
	err := a.handler.ExampleApp.DoSomethingInApp()
	if err != nil {
		return nil, err
	}
	return &example.GetTextMessageResponse{
		Data: nil,
	}, nil
}
