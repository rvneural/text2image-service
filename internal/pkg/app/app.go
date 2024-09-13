package app

import (
	handler "Text2ImageService/internal/api/rest"
	endpoint "Text2ImageService/internal/endpoint/app"
	service "Text2ImageService/internal/services"

	"github.com/rs/zerolog"
)

type App struct {
	endpoint *endpoint.App
	service  *service.Service
	handler  *handler.Handler
	logger   *zerolog.Logger
}

func New(logger *zerolog.Logger) *App {
	service := service.New(logger)
	handler := handler.NewHandler(service, logger)
	endpoint := endpoint.NewApp(handler, logger)

	return &App{endpoint: endpoint, service: service, handler: handler, logger: logger}
}

func (a *App) Run() error {
	a.logger.Info().Msg("Starting server...")
	return a.endpoint.Start()
}
