package docs

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func (a adapter) Swagger(ctx *fiber.Ctx) (err error) {
	url := a.config.App.Host
	swaggerUiUrl := fmt.Sprintf("%s/docs/swagger-ui", url)
	swaggerJsonUrl := fmt.Sprintf("%s/docs/swagger.yaml", url)

	err = ctx.Render("index", fiber.Map{
		"title":          "Example API",
		"swaggerUiUrl":   swaggerUiUrl,
		"swaggerJsonUrl": swaggerJsonUrl,
		"deepLinking":    a.config.Swagger.DeepLinking,
		"docExpansion":   a.config.Swagger.DocExpansion,
	})
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to render swagger-ui", slog.String("error", err.Error()))
	}

	return err
}
