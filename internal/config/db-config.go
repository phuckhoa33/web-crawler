package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// PostgresDBConfig struct
type PostgresDBConfig struct {
	PostgresDatabaseHost     string `env:"POSTGRES_DATABASE_HOST" required:"true"`
	PostgresDatabasePort     string `env:"POSTGRES_DATABASE_PORT" required:"true"`
	PostgresDatabaseName     string `env:"POSTGRES_DATABASE_NAME" required:"true"`
	PostgresDatabaseUser     string `env:"POSTGRES_DATABASE_USER" required:"true"`
	PostgresDatabasePassword string `env:"POSTGRES_DATABASE_PASSWORD" required:"true"`
	PostgresDatabaseDebug    string `env:"POSTGRES_DATABASE_DEBUG" required:"true"`
}

// DBConfig DbConfig struct
type DBConfig struct {
	PostgresConfig PostgresDBConfig
}

// LoadDBConfig function
func LoadDBConfig() DBConfig {
	var cfg DBConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
