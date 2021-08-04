package helper

import (
	"os"
)

func GetEnvWithDefault(key string, defaultValue string) (r string) {
	r = os.Getenv(key)
	if r == "" {
		r = defaultValue
	}
	return
}

