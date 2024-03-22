package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jrolstad/blog-monitor/internal/pkg/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/orchestrators"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

var (
	subscriptionRepository        repositories.SubscriptionRepository
	notificationHistoryRepository repositories.NotificationHistoryRepository
	notificationService           services.NotificationService
	secretService                 services.SecretService
)

func init() {
	config := configuration.NewAppConfig()
	subscriptionRepository = repositories.NewSubscriptionRepository(config)
	notificationHistoryRepository = repositories.NewNotificationHistoryRepository(config)
	notificationService = services.NewNotificationService(config)
	secretService = services.NewSecretService(config)

}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.CloudWatchEvent) error {

	logging.LogEvent("Lambda Event Handler", "action", "start")
	err := orchestrators.NotifyNewPosts(subscriptionRepository, notificationHistoryRepository, notificationService, secretService)
	logging.LogEvent("Lambda Event Handler", "action", "complete", "success", err == nil)
	return err
}
