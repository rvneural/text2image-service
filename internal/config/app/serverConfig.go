package app

import "os"

const (
	ADDR = ":8083"
)

var (
	BEARER_KEY = os.Getenv("KEY")
)
