package services

import (
	"errors"
	"os"
)

const Secret_GoogleApiKey string = "googleApiKey"

type SecretService interface {
	Get(name string) (string, error)
}

func NewSecretService() SecretService {
	return &ConfigSecretService{}
}

type ConfigSecretService struct {
}

func (*ConfigSecretService) Get(name string) (string, error) {
	switch name {
	case Secret_GoogleApiKey:
		return os.Getenv("blog_monitor_google_apikey"), nil
	default:
		return "", errors.New("secret not found")
	}
}
