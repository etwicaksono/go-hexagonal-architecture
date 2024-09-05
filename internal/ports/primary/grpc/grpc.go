package grpc

import (
	"github.com/etwicaksono/public-proto/gen/example"
)

type GrpcHandlerInterface interface {
	Run() error
	example.ExampleServiceServer
}
