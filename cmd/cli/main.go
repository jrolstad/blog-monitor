package main

import (
	"github.com/jrolstad/blog-monitor/internal/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/orchestrators"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

func main() {
	config := configuration.NewAppConfig()
	subscriptionRepository := repositories.NewSubscriptionRepository()
	notificationHistoryRepository := repositories.NewNotificationHistoryRepository()
	notificationService := services.NewNotificationService()

	err := orchestrators.NotifyNewPosts(config, subscriptionRepository, notificationHistoryRepository, notificationService)

	if err != nil {
		panic(err)
	}
}
