package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Env             string
	RestHost        string
	RestPort        int
	GrpcHost        string
	GrpcPort        int
	IdleTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	GracefulTimeout time.Duration
	Name            string
	Version         string
	Host            string
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
		panic(err.Error())
	}

	return Config{
		App: AppConfig{
			Env:             vpr.GetString("APP_ENV"),
			RestHost:        vpr.GetString("APP_REST_NATIVE_HOST"),
			RestPort:        vpr.GetInt("APP_REST_NATIVE_PORT"),
			GrpcHost:        vpr.GetString("APP_GRPC_HOST"),
			GrpcPort:        vpr.GetInt("APP_GRPC_PORT"),
			IdleTimeout:     vpr.GetDuration("APP_IDLE_TIMEOUT"),
			ReadTimeout:     vpr.GetDuration("APP_READ_TIMEOUT"),
			WriteTimeout:    vpr.GetDuration("APP_WRITE_TIMEOUT"),
			GracefulTimeout: vpr.GetDuration("APP_GRACEFUL_TIMEOUT"),
			Name:            vpr.GetString("APP_NAME"),
			Version:         vpr.GetString("APP_VERSION"),
			Host:            vpr.GetString("APP_HOST"),
		},
	}
}
