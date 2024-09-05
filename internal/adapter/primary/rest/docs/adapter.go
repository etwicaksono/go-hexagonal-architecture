package docs

import (
	"github.com/etwicaksono/go-hexagonal-architecture/config"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct {
	config config.Config
}

func NewDocumentationHandler(cfg config.Config) rest.SwaggerHandlerInterface {
	return &adapter{
		config: cfg,
	}
}
