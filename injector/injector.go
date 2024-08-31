//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(config.LoadConfig)

func LoggerInit() error {
	wire.Build(
		serverSet,
		loggerInit,
	)
	return nil
}

func RestProvider() *fiber.App {
	wire.Build(
		serverSet,
		rest.NewRestApp,
	)
	return nil
}
