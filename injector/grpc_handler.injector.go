package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
)

func grpcHandlerProvider(exampleApp app.ExampleAppInterface) grpc.Handler {
	return grpc.Handler{ExampleApp: exampleApp}
}
