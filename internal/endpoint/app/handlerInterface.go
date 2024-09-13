package app

import "github.com/labstack/echo/v4"

type Text2ImageHadler interface {
	HandleRequest(c echo.Context) error
}
