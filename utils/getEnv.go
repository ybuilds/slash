package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func GetValue(key string) string {
	val := os.Getenv(key)
	if val == "" {
		return ""
	}

	return val
}
