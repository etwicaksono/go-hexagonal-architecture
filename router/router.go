package router

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	docs rest.SwaggerHandlerInterface
}

func NewRouter(docs rest.SwaggerHandlerInterface) Router {
	return Router{
		docs: docs,
	}
}

func SetRoute(app *fiber.App, router Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/swagger/*", router.docs.Swagger)
}
