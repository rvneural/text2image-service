package app

import (
	handler "Text2ImageService/internal/api/rest"
	endpoint "Text2ImageService/internal/endpoint/app"
	service "Text2ImageService/internal/services"

	config "Text2ImageService/internal/config/app"
	db "Text2ImageService/internal/services/db"

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
	db := db.New(config.DB_URL)
	handler := handler.NewHandler(service, db, logger)
	endpoint := endpoint.NewApp(handler, logger)

	return &App{endpoint: endpoint, service: service, handler: handler, logger: logger}
}

func (a *App) Run() error {
	a.logger.Info().Msg("Starting server...")
	return a.endpoint.Start()
}
