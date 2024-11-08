package yandexart

import (
	"os"
	"time"
)

const (
	GENERATION_URL = "https://llm.api.cloud.yandex.net/foundationModels/v1/imageGenerationAsync"
	CHECK_URL      = "https://llm.api.cloud.yandex.net:443/operations/"
	MODEL_URI      = "art://b1gjtlqofdt5mu5io6a9/yandex-art/latest"
	DELAY_SEC      = 3 * time.Second
	WAIT_TIMEOUT   = 1 * time.Minute
)

var (
	API_KEY = os.Getenv("API_KEY")
)
