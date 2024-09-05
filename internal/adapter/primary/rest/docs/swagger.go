package docs

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (a adapter) Swagger(ctx *fiber.Ctx) (err error) {
	url := a.config.App.Host
	swaggerUiUrl := fmt.Sprintf("%s/docs/swagger-ui", url)
	swaggerJsonUrl := fmt.Sprintf("%s/docs/swagger.yaml", url)

	return ctx.Render("index", fiber.Map{
		"title":          "Example API",
		"swaggerUiUrl":   swaggerUiUrl,
		"swaggerJsonUrl": swaggerJsonUrl,
		"deepLinking":    a.config.Swagger.DeepLinking,
		"docExpansion":   a.config.Swagger.DocExpansion,
	})
}
