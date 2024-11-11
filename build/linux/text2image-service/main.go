package main

import (
	"Text2ImageService/cmd/log"
	"Text2ImageService/internal/pkg/app"
)

func main() {

	logger := log.NewLogger()

	app := app.New(&logger)
	logger.Fatal().Msg(app.Run().Error())
}
