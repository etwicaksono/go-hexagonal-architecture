//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest/docs"
	"github.com/etwicaksono/go-hexagonal-architecture/router"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var configSet = wire.NewSet(config.LoadConfig)
var routerSet = wire.NewSet(
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

func RestProvider() *fiber.App {
	wire.Build(
		routerSet,
		configSet,
		rest.NewRestApp,
	)
	return nil
}
