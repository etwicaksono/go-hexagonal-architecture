package example_app

import (
	"fmt"
	"log/slog"
	"time"
)

func (e exampleApp) DoSomethingInApp() error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	slog.InfoContext(e.ctx, fmt.Sprintf("Do something in app at: %s", currentTime))
	return e.core.DoSomethingInCore()
}
