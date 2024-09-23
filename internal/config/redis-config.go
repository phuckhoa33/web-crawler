package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type RedisConfig struct {
	RedisHost     string `env:"REDIS_HOST" required:"true"`
	RedisPort     string `env:"REDIS_PORT" required:"true"`
	RedisPassword string `env:"REDIS_PASSWORD" required:"false"`
	RedisDb       int    `env:"REDIS_DB" required:"false"`
}

func LoadRedisConfig() RedisConfig {
	var cfg RedisConfig

	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
