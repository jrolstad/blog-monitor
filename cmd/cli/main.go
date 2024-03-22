package main

import (
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/orchestrators"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

func main() {
	subscriptionRepository := repositories.NewSubscriptionRepository()
	notificationHistoryRepository := repositories.NewNotificationHistoryRepository()
	notificationService := services.NewNotificationService()
	secretService := services.NewSecretService()

	err := orchestrators.NotifyNewPosts(subscriptionRepository, notificationHistoryRepository, notificationService, secretService)

	if err != nil {
		logging.LogPanic(err)
	}
}
