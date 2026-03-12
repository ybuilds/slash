package utils

import (
	"os"
)

func GetValue(key string) string {
	val := os.Getenv(key)
	if val == "" {
		return ""
	}

	return val
}
