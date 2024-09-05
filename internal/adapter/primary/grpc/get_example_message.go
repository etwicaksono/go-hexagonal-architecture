package grpc

import (
	"context"
	"fmt"

	"github.com/etwicaksono/public-proto/gen/example"
)

func (a *adapter) GetExampleMessage(ctx context.Context, in *example.ExampleRequest) (*example.ExampleResponse, error) {
	err := a.handler.ExampleApp.DoSomethingInApp()
	if err != nil {
		return nil, err
	}
	return &example.ExampleResponse{
		Message: fmt.Sprintf("Here is message from grpc: %s", in.Message),
	}, nil
}
