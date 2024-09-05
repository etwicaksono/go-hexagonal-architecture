//go:build wireinject
// +build wireinject

package injector

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest/docs"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest/example_rest"
	"github.com/etwicaksono/go-hexagonal-architecture/router"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var configSet = wire.NewSet(config.LoadConfig)
var appSet = wire.NewSet(
	exampleAppConfigProvider,
	example_app.NewExampleApp,
)
var restSet = wire.NewSet(
	exampleRestConfigProvider,
	example_rest.NewExampleRestHandler,
)
var routerSet = wire.NewSet(
	restSet,
	docs.NewDocumentationHandler,
	router.NewRouter,
)

func LoggerInit() error {
	wire.Build(
		configSet,
		loggerInit,
	)
	return nil
}

func RestProvider(ctx context.Context) *fiber.App {
	wire.Build(
		routerSet,
		configSet,
		rest.NewRestApp,
	)
	return nil
}

func GrpcHandlerProvider() grpc.Handler {
	wire.Build(
		appSet,
		grpcHandlerProvider,
	)
	return grpc.Handler{}
}
