package config

import (
	"errors"
	"os"
)

func GetApiToken() (string, error) {
	token, exists := os.LookupEnv("TOKEN")

	if !exists {
		return "", errors.New("API Token not exists")
	}

	return token, nil
}
