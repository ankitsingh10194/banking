package main

import (
	"github.com/ankit/banking/app"
	"github.com/ankit/banking/logger"
)

func main() {

	logger.Info("Starting our application...")
	app.Start()
}
