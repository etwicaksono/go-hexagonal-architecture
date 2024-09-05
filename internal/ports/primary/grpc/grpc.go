package grpc

import (
	"context"
	"github.com/etwicaksono/public-proto/gen/example"
)

type ExampleGrpcHandlerInterface interface {
	Run() error

	/*
		ExampleGrpcHandlerInterface
	*/
	GetExample(ctx context.Context, in *example.ExampleRequest) (*example.ExampleResponse, error)
}
