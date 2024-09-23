package main

import (
	application "github.com/phuckhoa33/web-crawler"
	"github.com/phuckhoa33/web-crawler/internal/config"
)

// @title			Golang Boilerplate API
// @version		1.0
// @description	This is a boilerplate for building RESTful APIs in Golang.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// Init config
	config := config.NewConfig()

	// Start app
	application.Start(config)
}
