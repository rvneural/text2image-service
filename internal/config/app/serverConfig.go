package app

import "os"

const (
	ADDR = ":8083"
)

var (
	DB_URL = os.Getenv("DB_URL")
)
