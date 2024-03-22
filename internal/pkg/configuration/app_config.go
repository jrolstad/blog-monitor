package configuration

import "os"

type AppConfig struct {
	AwsRegion                    string
	EmailSender                  string
	SubscriptionTableName        string
	NotificationHistoryTableName string
	GoogleApiKeySecretName       string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AwsRegion:                    os.Getenv("aws_region"),
		EmailSender:                  os.Getenv("email_sender"),
		SubscriptionTableName:        os.Getenv("subscription_table_name"),
		NotificationHistoryTableName: os.Getenv("notification_history_table_name"),
		GoogleApiKeySecretName:       os.Getenv("secret_google_api_key"),
	}
}
