package router

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	docs    rest.SwaggerHandlerInterface
	example rest.ExampleHandlerInterface
}

func NewRouter(
	docs rest.SwaggerHandlerInterface,
	example rest.ExampleHandlerInterface,
) Router {
	return Router{
		docs:    docs,
		example: example,
	}
}

func SetRoute(app *fiber.App, router Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/swagger/*", router.docs.Swagger)

	// Example
	app.Get("/message/text", router.example.GetTextMessage)
	app.Post("/message/text", router.example.SendTextMessage)
	app.Get("/message/multimedia", router.example.GetMultimediaMessage)
	app.Post("/message/multimedia", router.example.SendMultimediaMessage)
}
