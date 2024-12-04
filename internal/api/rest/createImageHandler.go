package rest

import (
	"Text2ImageService/internal/models/json/client"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type DBWorker interface {
	RegisterOperation(uniqID string, operation_type string, user_id int) error
	SetResult(uniqID string, data []byte) error
}

type Handler struct {
	service  Service
	logger   *zerolog.Logger
	dbWorker DBWorker
}

func NewHandler(service Service, dbworker DBWorker, logger *zerolog.Logger) *Handler {
	return &Handler{service: service, logger: logger, dbWorker: dbworker}
}

func (h *Handler) HandleRequest(c echo.Context) error {
	h.logger.Info().Msg("New request received from server: " + c.RealIP())

	request := new(client.Request)
	err := c.Bind(request)
	if err != nil {
		h.logger.Error().Msg("Error binding request: " + err.Error())
		return c.JSON(http.StatusBadRequest, client.Error{Error: "Invalid request body", Details: err.Error()})
	}

	if request.Operation_ID != "" {
		id, err := strconv.Atoi(request.UserID)
		if err != nil {
			id = 0
		}
		go h.dbWorker.RegisterOperation(request.Operation_ID, "image", id)
	}

	b64Image, seed, err := h.service.ConvertTextToImage(request.Prompt, request.Seed,
		request.WidthRatio, request.HeightRatio)

	if err != nil {
		h.logger.Error().Msg("Error generating image: " + err.Error())
		return c.JSON(http.StatusInternalServerError, client.Error{Error: "Error generating image", Details: err.Error()})
	}

	var response client.Response
	response.Image.B64String = b64Image
	response.Image.Seed = seed

	if request.Operation_ID != "" {
		dbResult := client.DBResult{
			Prompt:    request.Prompt,
			Seed:      request.Seed,
			B64string: b64Image,
			Name:      "generated_image.jpg",
		}
		byteResponse, _ := json.Marshal(dbResult)
		go h.dbWorker.SetResult(request.Operation_ID, byteResponse)
	}

	return c.JSON(http.StatusOK, response)
}
