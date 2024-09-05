package example_core

import (
	"fmt"
	"log/slog"
	"time"
)

func (e exampleCore) DoSomethingInCore() error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	slog.InfoContext(e.ctx, fmt.Sprintf("Do something in core at: %s", currentTime))
	return nil
}
