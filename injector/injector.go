//go:build wireinject
// +build wireinject

package injector

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/authentication_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/authentication_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/authentication_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_message_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/middleware"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/auth_cache"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/minio"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"log/slog"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_message_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/example_message_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var configSet = wire.NewSet(config.LoadConfig)
var validatorSet = wire.NewSet(validatorProvider)
var routerSet = wire.NewSet(
	middleware.NewMiddleware,
	docs_handler.NewDocumentationHandler,
	rest.NewRouter,
)

var authenticationSet = wire.NewSet(
	auth_cache.NewCache,
	userDbProvider,
	rest_util.NewJwt,
	authentication_core.NewAuthenticationCore,
	authentication_app.NewAuthenticationApp,
	authentication_handler.NewAuthenticationRestHandler,
)
var exampleSet = wire.NewSet(
	minio.MinioProvider,
	validatorSet,
	messageDbProvider,
	example_message_core.NewExampleMessageCore,
	example_message_app.NewExampleMessageApp,
	example_message_handler.NewExampleRestHandler,
)

func LoggerInit() (logger *slog.Logger, err error) {
	wire.Build(
		configSet,
		loggerInit,
	)
	return nil, nil
}

func RestProvider(
	ctx context.Context,
	dbClient *entity.DbClient,
	redisClient *redis.Client,
) *fiber.App {
	wire.Build(
		configSet,
		routerSet,
		authenticationSet,
		exampleSet,
		rest.NewRestApp,
	)
	return nil
}

func GrpcHandlerProvider(
	ctx context.Context,
	dbClient *entity.DbClient,
) grpc.Handler {
	wire.Build(
		configSet,
		exampleSet,
		grpcHandlerProvider,
	)
	return grpc.Handler{}
}
