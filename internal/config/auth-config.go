package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AuthConfig struct {
	AccessTokenSecret     string `env:"ACCESS_TOKEN_SECRET" required:"true"`
	AccessTokenExpiredIn  string `env:"ACCESS_TOKEN_EXPIRES_IN" required:"true"`
	RefreshTokenSecret    string `env:"REFRESH_TOKEN_SECRET" required:"true"`
	RefreshTokenExpiredIn string `env:"REFRESH_TOKEN_EXPIRES_IN" required:"true"`
}

func LoadAuthConfig() AuthConfig {
	var cfg AuthConfig

	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
