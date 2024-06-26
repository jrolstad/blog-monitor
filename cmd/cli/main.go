package main

import (
	"github.com/jrolstad/blog-monitor/internal/pkg/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/orchestrators"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

func main() {
	config := configuration.NewAppConfig()
	subscriptionRepository := repositories.NewSubscriptionRepository(config)
	notificationHistoryRepository := repositories.NewNotificationHistoryRepository(config)
	notificationService := services.NewNotificationService(config)
	secretService := services.NewSecretService(config)

	err := orchestrators.NotifyNewPosts(subscriptionRepository, notificationHistoryRepository, notificationService, secretService)

	if err != nil {
		logging.LogPanic(err)
	}
}
