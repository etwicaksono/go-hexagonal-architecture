package rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/authentication_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/docs_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/example_message_handler"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/framework/primary/rest/middleware"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/rest_util"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	middleware     *middleware.Middleware
	docs           docs_handler.DocsHandler
	authentication authentication_handler.AuthenticationHandler
	exampleMessage example_message_handler.ExampleMessageHandler
}

func NewRouter(
	middleware *middleware.Middleware,
	docs docs_handler.DocsHandler,
	authentication authentication_handler.AuthenticationHandler,
	exampleMessage example_message_handler.ExampleMessageHandler,
) Router {
	return Router{
		middleware:     middleware,
		docs:           docs,
		authentication: authentication,
		exampleMessage: exampleMessage,
	}
}

func SetRoute(app *fiber.App, router Router) {
	jwtMiddleware := router.middleware.Jwt.JwtAuthenticate

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
	example := app.Group("/example", jwtMiddleware)
	example.Get("/message/text", router.exampleMessage.GetTextMessage)
	example.Post("/message/text", router.exampleMessage.SendTextMessage)
	example.Get("/message/multimedia", router.exampleMessage.GetMultimediaMessage)
	example.Post("/message/multimedia", router.exampleMessage.SendMultimediaMessage)
}
