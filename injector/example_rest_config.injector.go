package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/primary/rest/example_rest"
)

func exampleRestConfigProvider() example_rest.Config {
	return example_rest.Config{}
}
