package yandexart

import "time"

const (
	GENERATION_URL = "https://llm.api.cloud.yandex.net/foundationModels/v1/imageGenerationAsync"
	CHECK_URL      = "https://llm.api.cloud.yandex.net:443/operations/"
	API_KEY        = "AQVNw7lDrwxeeWvbXv9CRPiGCam43_hHCITPxRqp"
	MODEL_URI      = "art://b1gjtlqofdt5mu5io6a9/yandex-art/latest"
	DELAY_SEC      = 3 * time.Second
	WAIT_TIMEOUT   = 1 * time.Minute
)
