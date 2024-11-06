package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/config"
	"log/slog"
	"os"
)

func loggerInit(cfg config.Config) error {
	var logLevel slog.Level

	switch cfg.App.LogLevel {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	}

	//Initiate logger
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.String("service", cfg.App.Name),
		slog.String("with-release", cfg.App.Version),
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	return nil
}
