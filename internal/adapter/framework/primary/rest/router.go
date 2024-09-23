package rest

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/rest_util"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	docs           rest.SwaggerHandlerInterface
	exampleMessage rest.ExampleMessageHandlerInterface
}

func NewRouter(
	docs rest.SwaggerHandlerInterface,
	exampleMessage rest.ExampleMessageHandlerInterface,
) Router {
	return Router{
		docs:           docs,
		exampleMessage: exampleMessage,
	}
}

func SetRoute(app *fiber.App, router Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return rest_util.ResponseOk(c, "Welcome to Hexagonal Architecture!")
	})
	app.Get("/swagger/*", router.docs.Swagger)

	// Example Message
	app.Get("/message/text", router.exampleMessage.GetTextMessage)
	app.Post("/message/text", router.exampleMessage.SendTextMessage)
	app.Get("/message/multimedia", router.exampleMessage.GetMultimediaMessage)
	app.Post("/message/multimedia", router.exampleMessage.SendMultimediaMessage)
}
