package configuration

import "os"

type AppConfig struct {
	AwsRegion   string
	EmailSender string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AwsRegion:   os.Getenv("aws_region"),
		EmailSender: os.Getenv("email_sender"),
	}
}
