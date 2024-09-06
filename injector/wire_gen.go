// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/example_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/example_mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/infrastructure"
	"github.com/etwicaksono/go-hexagonal-architecture/router"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// Injectors from injector.go:

func LoggerInit() error {
	configConfig := config.LoadConfig()
	error2 := loggerInit(configConfig)
	return error2
}

func RestProvider(ctx context.Context) *fiber.App {
	configConfig := config.LoadConfig()
	swaggerHandlerInterface := docs.NewDocumentationHandler(ctx, configConfig)
	mongoInterface := infrastructure.NewMongo(ctx, configConfig)
	exampleDbInterface := example_mongo.NewExampleMongo(configConfig, mongoInterface)
	exampleCoreInterface := example_core.NewExampleCore(ctx, exampleDbInterface)
	exampleAppInterface := example_app.NewExampleApp(ctx, exampleCoreInterface)
	exampleHandlerInterface := example_rest.NewExampleRestHandler(exampleAppInterface)
	routerRouter := router.NewRouter(swaggerHandlerInterface, exampleHandlerInterface)
	app := rest.NewRestApp(ctx, configConfig, routerRouter)
	return app
}

func GrpcHandlerProvider(ctx context.Context) grpc.Handler {
	configConfig := config.LoadConfig()
	mongoInterface := infrastructure.NewMongo(ctx, configConfig)
	exampleDbInterface := example_mongo.NewExampleMongo(configConfig, mongoInterface)
	exampleCoreInterface := example_core.NewExampleCore(ctx, exampleDbInterface)
	exampleAppInterface := example_app.NewExampleApp(ctx, exampleCoreInterface)
	handler := grpcHandlerProvider(exampleAppInterface)
	return handler
}

// injector.go:

var configSet = wire.NewSet(config.LoadConfig)

var exampleSet = wire.NewSet(
	configSet, infrastructure.NewMongo, example_mongo.NewExampleMongo, example_core.NewExampleCore, example_app.NewExampleApp,
)

var routerSet = wire.NewSet(example_rest.NewExampleRestHandler, docs.NewDocumentationHandler, router.NewRouter)
