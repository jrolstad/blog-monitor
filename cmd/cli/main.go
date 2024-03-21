package main

import (
	"github.com/jrolstad/blog-monitor/internal/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/orchestrators"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
)

func main() {
	config := configuration.NewAppConfig()
	subscriptionRepository := repositories.NewSubscriptionRepository()

	err := orchestrators.NotifyNewPosts(config, subscriptionRepository)

	if err != nil {
		panic(err)
	}
}
