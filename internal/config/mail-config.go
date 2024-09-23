package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type MailConfig struct {
	MailHost     string `env:"MAIL_HOST" required:"true"`
	MailPort     int    `env:"MAIL_PORT" required:"true"`
	MailUsername string `env:"MAIL_SMTP_USERNAME" required:"true"`
	MailPassword string `env:"MAIL_SMTP_PASSWORD" required:"true"`
	MailSecure   bool   `env:"MAIL_SMTP_SECURE" required:"true"`
}

func LoadMailConfig() MailConfig {

	var cfg MailConfig

	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
