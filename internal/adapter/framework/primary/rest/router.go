package rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	docs           rest.SwaggerHandlerInterface
	authentication rest.AuthenticationHandlerInterface
	exampleMessage rest.ExampleMessageHandlerInterface
}

func NewRouter(
	docs rest.SwaggerHandlerInterface,
	authentication rest.AuthenticationHandlerInterface,
	exampleMessage rest.ExampleMessageHandlerInterface,
) Router {
	return Router{
		docs:           docs,
		authentication: authentication,
		exampleMessage: exampleMessage,
	}
}

func SetRoute(app *fiber.App, router Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return rest_util.ResponseOk(c, "Welcome to Hexagonal Architecture!")
	})
	app.Get("/swagger/*", router.docs.Swagger)

	// Authentication
	auth := app.Group("/auth")
	auth.Post("/register", router.authentication.Register)
	auth.Post("/login", router.authentication.Login)
	auth.Post("/logout", router.authentication.Logout)
	auth.Post("/refresh", router.authentication.Refresh)

	// Example
	example := app.Group("/example")
	example.Get("/message/text", router.exampleMessage.GetTextMessage)
	example.Post("/message/text", router.exampleMessage.SendTextMessage)
	example.Get("/message/multimedia", router.exampleMessage.GetMultimediaMessage)
	example.Post("/message/multimedia", router.exampleMessage.SendMultimediaMessage)
}
