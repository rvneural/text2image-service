package services

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"

	config "Text2ImageService/internal/config/yandexart"
	yandexart "Text2ImageService/internal/models/json/yandexart"
)

func (s *Service) verifyParams(seed *string, widthRatio *string, heightRatio *string) {
	if *widthRatio == "" {
		*widthRatio = "3"
	}

	if *heightRatio == "" {
		*heightRatio = "2"
	}

	if strings.EqualFold(*seed, "") || strings.EqualFold(*seed, "random") {
		var randomNumber int64 = rand.Int64()
		*seed = strconv.FormatInt(randomNumber, 10)
		s.logger.Debug().Msg("Generated random seed: " + *seed)
	}
}

func (s *Service) initRequestToYandexArt(prompt, seed, widthRatio, heightRatio string) yandexart.Request {
	var Request yandexart.Request

	Request.ModelUri = config.MODEL_URI
	Request.GenerationOptions.Seed = seed
	Request.GenerationOptions.AspectRatio.WidthRatio = widthRatio
	Request.GenerationOptions.AspectRatio.HeightRatio = heightRatio
	Request.Messages = append(Request.Messages, yandexart.Message{Weight: "1.0", Text: prompt})

	return Request
}

func (s *Service) waitForReadyStatus(operationId string) (bool, error) {
	httpRequest, err := http.NewRequest("GET", config.CHECK_URL+operationId, nil)
	if err != nil {
		return false, err
	}
	httpRequest.Header.Set("Authorization", "Api-Key "+config.API_KEY)
	client := &http.Client{}

	for {
		var sendResponse yandexart.Response
		<-time.After(config.DELAY_SEC)
		httpResponse, err := client.Do(httpRequest)

		if err != nil {
			return false, err
		}

		defer httpResponse.Body.Close()

		byteResponse, err := io.ReadAll(httpResponse.Body)

		if err != nil {
			return false, err
		}

		err = json.Unmarshal(byteResponse, &sendResponse)

		if err != nil {
			return false, err
		}

		if sendResponse.Done {
			return true, nil
		}
	}
}

func (s *Service) getResultImage(operationId string) (string, error) {

	done, err := s.waitForReadyStatus(operationId)

	if err != nil {
		return "", err
	}
	if !done {
		return "", errors.New("operation not ready")
	}

	httpRequest, err := http.NewRequest("GET", config.CHECK_URL+operationId, nil)
	if err != nil {
		return "", err
	}
	httpRequest.Header.Set("Authorization", "Api-Key "+config.API_KEY)
	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer httpResponse.Body.Close()

	var Response yandexart.Response
	byteResponse, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(byteResponse, &Response)
	if err != nil {
		return "", err
	}

	b64Image := Response.Response.Image

	return b64Image, nil
}
