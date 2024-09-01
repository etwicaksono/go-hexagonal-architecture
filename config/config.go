package config

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/spf13/viper"
	"log/slog"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
	App     AppConfig
	Swagger SwaggerConfig
}

type AppConfig struct {
	Env                  string
	RestHost             string
	RestPort             int
	RestPrefork          bool
	RestRecovery         bool
	RestEnableStackTrace bool
	RestCorsAllowOrigins string
	RestCorsAllowHeaders string
	RestCorsAllowMethods string
	GrpcHost             string
	GrpcPort             int
	IdleTimeout          time.Duration
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	GracefulTimeout      time.Duration
	Name                 string
	Version              string
	Host                 string
}

type SwaggerConfig struct {
	DeepLinking  bool
	DocExpansion string
}

func LoadConfig() Config {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	projectRoot := filepath.Join(filepath.Dir(b), "../")
	vpr := viper.New()

	vpr.AddConfigPath(projectRoot)
	vpr.SetConfigFile(".env")
	vpr.AutomaticEnv()

	err := vpr.ReadInConfig()
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to read config file", slog.String(entity.Error, err.Error()))
		panic(err.Error())
	}

	return Config{
		App: AppConfig{
			Env:                  vpr.GetString("APP_ENV"),
			RestHost:             vpr.GetString("APP_REST_HOST"),
			RestPort:             vpr.GetInt("APP_REST_PORT"),
			RestPrefork:          vpr.GetBool("FIBER_PREFORK"),
			RestRecovery:         vpr.GetBool("FIBER_RECOVERY"),
			RestEnableStackTrace: vpr.GetBool("FIBER_ENABLE_STACK_TRACE"),
			RestCorsAllowOrigins: vpr.GetString("FIBER_CORS_ALLOW_ORIGINS"),
			RestCorsAllowHeaders: vpr.GetString("FIBER_CORS_ALLOW_HEADERS"),
			RestCorsAllowMethods: vpr.GetString("FIBER_CORS_ALLOW_METHODS"),
			GrpcHost:             vpr.GetString("APP_GRPC_HOST"),
			GrpcPort:             vpr.GetInt("APP_GRPC_PORT"),
			IdleTimeout:          vpr.GetDuration("APP_IDLE_TIMEOUT"),
			ReadTimeout:          vpr.GetDuration("APP_READ_TIMEOUT"),
			WriteTimeout:         vpr.GetDuration("APP_WRITE_TIMEOUT"),
			GracefulTimeout:      vpr.GetDuration("APP_GRACEFUL_TIMEOUT"),
			Name:                 vpr.GetString("APP_NAME"),
			Version:              vpr.GetString("APP_VERSION"),
			Host:                 vpr.GetString("APP_HOST"),
		},
		Swagger: SwaggerConfig{
			DeepLinking:  vpr.GetBool("SWAGGER_DEEP_LINKING"),
			DocExpansion: vpr.GetString("SWAGGER_DOC_EXPANSION"),
		},
	}
}
