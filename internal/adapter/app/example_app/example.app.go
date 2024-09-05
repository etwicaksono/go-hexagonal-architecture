package example_app

import "github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"

type exampleApp struct {
}

func (e exampleApp) DoSomething() error {
	//TODO implement me
	panic("implement me")
}

type Config struct {
}

func NewExampleApp(config Config) app.ExampleAppInterface {
	return &exampleApp{}
}
