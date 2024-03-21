package configuration

import "os"

type AppConfig struct {
	GoogleApiKey string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		GoogleApiKey: os.Getenv("blog_monitor_google_apikey"),
	}
}
