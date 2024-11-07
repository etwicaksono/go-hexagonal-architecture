//go:build wireinject
// +build wireinject

package injector

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/authentication_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/authentication_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/authentication_rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_message_rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/minio"
	mongo2 "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/user_mongo"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/mongo/example_message_mongo"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_message_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/example_message_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var configSet = wire.NewSet(config.LoadConfig)
var validatorSet = wire.NewSet(validatorProvider)
var routerSet = wire.NewSet(
	example_message_rest.NewExampleRestHandler,
	docs.NewDocumentationHandler,
	rest.NewRouter,
)

var authenticationSet = wire.NewSet(
	user_mongo.NewUserMongo,
	authentication_core.NewAuthenticationCore,
	authentication_app.NewAuthenticationApp,
	authentication_rest.NewAuthenticationRestHandler,
)
var exampleSet = wire.NewSet(
	configSet,
	minio.MinioProvider,
	validatorSet,
	mongo2.NewMongo,
	example_message_mongo.NewExampleMessageMongo,
	example_message_app.NewExampleMessageApp,
	example_message_core.NewExampleMessageCore,
)

func LoggerInit() error {
	wire.Build(
		configSet,
		loggerInit,
	)
	return nil
}

func RestProvider(
	ctx context.Context,
	mongoClient *mongo.Client,
) *fiber.App {
	wire.Build(
		routerSet,
		authenticationSet,
		exampleSet,
		rest.NewRestApp,
	)
	return nil
}

func GrpcHandlerProvider(
	ctx context.Context,
	mongoClient *mongo.Client,
) grpc.Handler {
	wire.Build(
		exampleSet,
		grpcHandlerProvider,
	)
	return grpc.Handler{}
}
