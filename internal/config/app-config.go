package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	SwaggerEnable  string `env:"SWAGGER_ENABLE" required:"true"`
	AppName        string `env:"APP_NAME" required:"true"`
	AppDescription string `env:"APP_DESCRIPTION" required:"true"`
	AppVersion     string `env:"APP_VERSION" required:"true"`
	AppEnv         string `env:"APP_ENV" required:"true"`
	AppPort        string `env:"APP_PORT" required:"true"`
	AppHost        string `env:"APP_HOST" required:"true"`
	AppApiDebug    string `env:"APP_API_DEBUG" required:"true"`
	AppApiPrefix   string `env:"APP_API_PREFIX" required:"true"`
}

func LoadAppConfig() AppConfig {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}
