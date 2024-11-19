package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	config "Text2ImageService/internal/config/app"

	"github.com/rs/zerolog"
)

type App struct {
	Text2ImageHadler Text2ImageHadler
	logger           *zerolog.Logger
}

func NewApp(text2ImageHadler Text2ImageHadler, logger *zerolog.Logger) *App {
	return &App{Text2ImageHadler: text2ImageHadler, logger: logger}
}

func (a *App) Start() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	a.logger.Info().Msg("Starting server on " + config.ADDR)

	e.POST("/", a.Text2ImageHadler.HandleRequest)
	return e.Start(config.ADDR)
}
