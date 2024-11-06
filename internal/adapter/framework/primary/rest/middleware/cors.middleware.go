package middleware

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware(app *fiber.App, config config.AppConfig) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.RestCorsAllowOrigins,
		AllowHeaders: config.RestCorsAllowHeaders,
		AllowMethods: config.RestCorsAllowMethods,
	}))
}
