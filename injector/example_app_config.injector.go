package injector

import "github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/app/example_app"

func exampleAppConfigProvider() example_app.Config {
	return example_app.Config{}
}
