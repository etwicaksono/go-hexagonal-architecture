package example_app

func (e exampleApp) DoSomethingInApp() error {
	return e.core.DoSomethingInCore()
}
