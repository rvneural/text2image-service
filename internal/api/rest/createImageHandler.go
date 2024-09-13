package rest

import (
	client2 "Text2ImageService/internal/models/json/client"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Handler struct {
	service Service
	logger  *zerolog.Logger
}

func NewHandler(service Service, logger *zerolog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) HandleRequest(c echo.Context) error {
	h.logger.Info().Msg("New request received from server: " + c.RealIP())

	request := new(client2.Request)
	err := c.Bind(request)

	if err != nil {
		h.logger.Error().Msg("Error binding request: " + err.Error())
		return c.JSON(http.StatusBadRequest, client2.Error{Error: "Invalid request body", Details: err.Error()})
	}

	b64Image, seed, err := h.service.ConvertTextToImage(request.Prompt, request.Seed,
		request.WidthRatio, request.HeightRatio)

	if err != nil {
		h.logger.Error().Msg("Error generating image: " + err.Error())
		return c.JSON(http.StatusInternalServerError, client2.Error{Error: "Error generating image", Details: err.Error()})
	}

	var response client2.Response
	response.Image.B64String = b64Image
	response.Image.Seed = seed

	return c.JSON(http.StatusOK, response)
}
