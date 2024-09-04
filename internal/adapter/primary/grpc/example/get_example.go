package example

import (
	"context"
	"github.com/etwicaksono/public-proto/gen/example"
)

func (a *adapter) GetExample(ctx context.Context, in *example.ExampleRequest) (*example.ExampleResponse, error) {
	return &example.ExampleResponse{
		Message: "Hello World",
	}, nil
}
