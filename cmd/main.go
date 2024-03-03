package main

import (
	app "ingressPlugin/internal/app"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	application := app.New(logger)
	err := application.GRPC.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
