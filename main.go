package main

import (
	"github.com/ankitsingh10194/banking/app"
	"github.com/ankitsingh10194/banking/logger"
)

func main() {

	logger.Info("Starting our application...")
	app.Start()
}
