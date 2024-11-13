package docs_handler

import (
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	errorsConst "github.com/etwicaksono/go-hexagonal-architecture/internal/errors"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (a DocsHandler) SwaggerPage(ctx *fiber.Ctx) (err error) {
	url := a.config.App.Host
	swaggerUiUrl := fmt.Sprintf("%s/docs/swagger-ui", url)
	swaggerJsonUrl := fmt.Sprintf("%s/swagger.yaml", url)

	err = ctx.Render("swagger-ui/index", fiber.Map{
		"title":          "Example API",
		"swaggerUiUrl":   swaggerUiUrl,
		"swaggerJsonUrl": swaggerJsonUrl,
		"deepLinking":    a.config.Swagger.DeepLinking,
		"docExpansion":   a.config.Swagger.DocExpansion,
	})
	if err != nil {
		slog.ErrorContext(a.ctx, "Failed to render swagger-ui", slog.String(entity.Error, err.Error()))
	}

	return err
}
func (a DocsHandler) SwaggerYaml(ctx *fiber.Ctx) (err error) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tmpl, err := template.ParseFiles(filepath.Join(wd, "/docs/swagger.yaml")) // Using a template file
	if err != nil {
		slog.ErrorContext(ctx.UserContext(), "Failed to parse template", slog.String(entity.Error, err.Error()))
		return errorsConst.ErrInternalServer
	}

	// Define dynamic data for your servers
	data := struct {
		ServerUrl         string
		ServerDescription string
	}{
		ServerUrl:         a.config.App.Host, // Use environment variable or config for dynamic host
		ServerDescription: "Local server",    // Define your base path if needed
	}

	// Render the template with the data
	var renderedSwagger strings.Builder
	if err := tmpl.Execute(&renderedSwagger, data); err != nil {
		slog.ErrorContext(ctx.UserContext(), "Failed to render template", slog.String(entity.Error, err.Error()))
		return errorsConst.ErrInternalServer
	}

	return ctx.Type("text/yaml").SendString(renderedSwagger.String())
}
