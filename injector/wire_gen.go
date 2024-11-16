// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/authentication_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_message_app"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/authentication_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/example_message_core"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/grpc"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/authentication_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_message_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/middleware"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/cache/auth_cache"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/secondary/minio"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

// Injectors from injector.go:

func LoggerInit() (*slog.Logger, error) {
	configConfig := config.LoadConfig()
	logger, err := loggerInit(configConfig)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func RestProvider(ctx context.Context, dbClient *entity.DbClient, redisClient *redis.Client) *fiber.App {
	configConfig := config.LoadConfig()
	authCacheInterface := auth_cache.NewCache(redisClient)
	jwt := rest_util.NewJwt(configConfig, authCacheInterface)
	middlewareMiddleware := middleware.NewMiddleware(jwt)
	docsHandler := docs_handler.NewDocumentationHandler(ctx, configConfig)
	userDbInterface := userDbProvider(configConfig, dbClient)
	authenticationCoreInterface := authentication_core.NewAuthenticationCore(userDbInterface, configConfig, jwt, authCacheInterface)
	validate := validatorProvider()
	authenticationAppInterface := authentication_app.NewAuthenticationApp(authenticationCoreInterface, validate, jwt)
	authenticationHandler := authentication_handler.NewAuthenticationRestHandler(authenticationAppInterface, jwt)
	exampleMessageDbInterface := messageDbProvider(configConfig, dbClient)
	minioInterface := minio.MinioProvider(ctx, configConfig)
	exampleMessageCoreInterface := example_message_core.NewExampleMessageCore(exampleMessageDbInterface, minioInterface, configConfig)
	exampleMessageAppInterface := example_message_app.NewExampleMessageApp(exampleMessageCoreInterface, validate)
	exampleMessageHandler := example_message_handler.NewExampleRestHandler(exampleMessageAppInterface, jwt)
	router := rest.NewRouter(middlewareMiddleware, docsHandler, authenticationHandler, exampleMessageHandler)
	app := rest.NewRestApp(configConfig, router)
	return app
}

func GrpcHandlerProvider(ctx context.Context, dbClient *entity.DbClient) grpc.Handler {
	configConfig := config.LoadConfig()
	exampleMessageDbInterface := messageDbProvider(configConfig, dbClient)
	minioInterface := minio.MinioProvider(ctx, configConfig)
	exampleMessageCoreInterface := example_message_core.NewExampleMessageCore(exampleMessageDbInterface, minioInterface, configConfig)
	validate := validatorProvider()
	exampleMessageAppInterface := example_message_app.NewExampleMessageApp(exampleMessageCoreInterface, validate)
	handler := grpcHandlerProvider(exampleMessageAppInterface)
	return handler
}

// injector.go:

var configSet = wire.NewSet(config.LoadConfig)

var validatorSet = wire.NewSet(validatorProvider)

var routerSet = wire.NewSet(middleware.NewMiddleware, docs_handler.NewDocumentationHandler, rest.NewRouter)

var authenticationSet = wire.NewSet(auth_cache.NewCache, userDbProvider, rest_util.NewJwt, authentication_core.NewAuthenticationCore, authentication_app.NewAuthenticationApp, authentication_handler.NewAuthenticationRestHandler)

var exampleSet = wire.NewSet(minio.MinioProvider, validatorSet,
	messageDbProvider, example_message_core.NewExampleMessageCore, example_message_app.NewExampleMessageApp, example_message_handler.NewExampleRestHandler,
)
