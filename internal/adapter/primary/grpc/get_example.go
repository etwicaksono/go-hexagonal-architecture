package grpc

import (
	"context"
	"github.com/etwicaksono/public-proto/gen/example"
)

func (a *adapter) GetExample(ctx context.Context, in *example.ExampleRequest) (*example.ExampleResponse, error) {
	err := a.handler.exampleApp.DoSomething()
	if err != nil {
		return nil, err
	}
	return &example.ExampleResponse{
		Message: "Hello World",
	}, nil
}
