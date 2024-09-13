package services

import (
	config "Text2ImageService/internal/config/yandexart"
	"Text2ImageService/internal/models/json/yandexart"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

type Service struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) ConvertTextToImage(prompt, seed, widthRatio, heightRatio string) (string, string, error) {

	logLine := fmt.Sprintf("Converting text to image\nPrompt: %s\nSeed: %s\nWidthRatio-HeightRatio: %s-%s", prompt, seed, widthRatio, heightRatio)
	s.logger.Info().Msg(logLine)

	s.verifyParams(&seed, &widthRatio, &heightRatio)

	var Request = s.initRequestToYandexArt(prompt, seed, widthRatio, heightRatio)

	seed = Request.GenerationOptions.Seed

	byteRequest, err := json.Marshal(Request)
	if err != nil {
		s.logger.Error().Msg("Error marshaling request JSON: " + err.Error())
		return "", "", err
	}

	// Отправка запроса на генерацию изображения
	httpRequest, err := http.NewRequest("POST", config.GENERATION_URL, bytes.NewBuffer(byteRequest))

	if err != nil {
		s.logger.Error().Msg("Error creating HTTP request: " + err.Error())
		return "", "", err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", "Api-Key "+config.API_KEY)

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		s.logger.Error().Msg("Error sending HTTP request: " + err.Error())
		return "", "", err
	}
	defer httpResponse.Body.Close()

	// Получение идентификатора операции
	byteResponse, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		s.logger.Error().Msg("Error reading response body: " + err.Error())
		return "", "", err
	}

	var sendResponse yandexart.Response

	err = json.Unmarshal(byteResponse, &sendResponse)

	if err != nil {
		s.logger.Error().Msg("Error unmarshalling send response JSON: " + err.Error())
		return "", "", err
	}

	if sendResponse.Error != "" {
		s.logger.Error().Msg("Error from Yandex ART: " + sendResponse.Error)
		return "", "", errors.New(sendResponse.Error)
	}

	operationId := sendResponse.Id

	b64Image, err := s.getResultImage(operationId)

	if err != nil {
		s.logger.Error().Msg("Error getting image from Yandex ART: " + err.Error())
		return "", "", err
	}

	return b64Image, seed, nil
}
