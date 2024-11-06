package docs

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"

	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/rest"
)

type adapter struct {
	ctx    context.Context
	config config.Config
}

func NewDocumentationHandler(
	ctx context.Context,
	cfg config.Config,
) rest.SwaggerHandlerInterface {
	return &adapter{
		ctx:    ctx,
		config: cfg,
	}
}
