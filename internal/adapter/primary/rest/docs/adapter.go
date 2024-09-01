package docs

import (
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
	"github.com/gofiber/fiber/v2"
)

type adapter struct {
	config config.Config
}

func NewDocumentationHandler(cfg config.Config) rest.SwaggerHandlerInterface {
	return &adapter{
		config: cfg,
	}
}

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
