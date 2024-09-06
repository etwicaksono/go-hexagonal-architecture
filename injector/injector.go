//go:build wireinject
// +build wireinject

package injector

import (
	"context"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/infrastructure"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/example_mongo"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_rest"

	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/example_core"
	"github.com/etwicaksono/go-hexagonal-architecture/router"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var configSet = wire.NewSet(config.LoadConfig)
var exampleSet = wire.NewSet(
	configSet,
	infrastructure.NewMongo,
	example_mongo.NewExampleMongo,
	example_core.NewExampleCore,
	example_app.NewExampleApp,
)
var routerSet = wire.NewSet(
	example_rest.NewExampleRestHandler,
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
		exampleSet,
		routerSet,
		rest.NewRestApp,
	)
	return nil
}

func GrpcHandlerProvider(ctx context.Context) grpc.Handler {
	wire.Build(
		exampleSet,
		grpcHandlerProvider,
	)
	return grpc.Handler{}
}
