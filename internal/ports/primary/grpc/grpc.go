package grpc

import (
	"context"

	"github.com/etwicaksono/public-proto/gen/example"
)

type GrpcHandlerInterface interface {
	Run() error

	/*
		ExampleGrpcHandlerInterface
	*/
	GetExampleMessage(ctx context.Context, in *example.ExampleRequest) (*example.ExampleResponse, error)
}
