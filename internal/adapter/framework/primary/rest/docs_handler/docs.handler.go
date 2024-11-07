package docs_handler

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
)

type DocsHandler struct {
	ctx    context.Context
	config config.Config
}

func NewDocumentationHandler(
	ctx context.Context,
	cfg config.Config,
) DocsHandler {
	return DocsHandler{
		ctx:    ctx,
		config: cfg,
	}
}
