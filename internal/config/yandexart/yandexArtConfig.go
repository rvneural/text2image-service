package yandexart

import (
	"os"
	"time"
)

const (
	GENERATION_URL = "https://llm.api.cloud.yandex.net/foundationModels/v1/imageGenerationAsync"
	CHECK_URL      = "https://llm.api.cloud.yandex.net:443/operations/"
	DELAY_SEC      = 3 * time.Second
	WAIT_TIMEOUT   = 1 * time.Minute
)

var (
	API_KEY   = os.Getenv("API_KEY")
	MODEL_URI = "art://" + os.Getenv("STORAGE_ID") + "/yandex-art/latest"
)
