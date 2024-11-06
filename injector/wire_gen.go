// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_message_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/example_message_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/authentication_rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_message_rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/minio"
	mongo2 "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/example_message_mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from injector.go:

func LoggerInit() error {
	configConfig := config.LoadConfig()
	error2 := loggerInit(configConfig)
	return error2
}

func RestProvider(ctx context.Context, mongoClient *mongo.Client) *fiber.App {
	configConfig := config.LoadConfig()
	swaggerHandlerInterface := docs.NewDocumentationHandler(ctx, configConfig)
	authenticationHandlerInterface := authentication_rest.NewAuthenticationRestHandler()
	exampleMessageDbInterface := example_message_mongo.NewExampleMessageMongo(configConfig, mongoClient)
	minioInterface := minio.MinioProvider(ctx, configConfig)
	exampleMessageCoreInterface := example_message_core.NewExampleMessageCore(exampleMessageDbInterface, minioInterface)
	validate := validatorProvider()
	exampleMessageAppInterface := example_message_app.NewExampleMessageApp(exampleMessageCoreInterface, validate)
	exampleMessageHandlerInterface := example_message_rest.NewExampleRestHandler(exampleMessageAppInterface)
	router := rest.NewRouter(swaggerHandlerInterface, authenticationHandlerInterface, exampleMessageHandlerInterface)
	app := rest.NewRestApp(ctx, configConfig, router)
	return app
}

func GrpcHandlerProvider(ctx context.Context, mongoClient *mongo.Client) grpc.Handler {
	configConfig := config.LoadConfig()
	exampleMessageDbInterface := example_message_mongo.NewExampleMessageMongo(configConfig, mongoClient)
	minioInterface := minio.MinioProvider(ctx, configConfig)
	exampleMessageCoreInterface := example_message_core.NewExampleMessageCore(exampleMessageDbInterface, minioInterface)
	validate := validatorProvider()
	exampleMessageAppInterface := example_message_app.NewExampleMessageApp(exampleMessageCoreInterface, validate)
	handler := grpcHandlerProvider(exampleMessageAppInterface)
	return handler
}

// injector.go:

var configSet = wire.NewSet(config.LoadConfig)

var validatorSet = wire.NewSet(validatorProvider)

var routerSet = wire.NewSet(example_message_rest.NewExampleRestHandler, docs.NewDocumentationHandler, rest.NewRouter)

var authenticationSet = wire.NewSet(authentication_rest.NewAuthenticationRestHandler)

var exampleSet = wire.NewSet(
	configSet, minio.MinioProvider, validatorSet, mongo2.NewMongo, example_message_mongo.NewExampleMessageMongo, example_message_app.NewExampleMessageApp, example_message_core.NewExampleMessageCore,
)
