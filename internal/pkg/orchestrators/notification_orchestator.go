package orchestrators

import (
	"errors"
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/processors"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

func NotifyNewPosts(subscriptionRepository repositories.SubscriptionRepository,
	notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService,
	secretService services.SecretService) error {

	subscriptions, err := subscriptionRepository.GetSubscriptions()
	if err != nil {
		return err
	}

	logging.LogEvent("NotifyNewPosts", "subscriptions", len(subscriptions), "status", "started")

	processingErrors := make([]error, 0)
	for _, item := range subscriptions {
		processor, err := processors.NewSubscriptionProcessor(item, secretService)
		err = processor.ProcessSubscription(notificationHistoryRepository, notificationService, item)
		if err != nil {
			processingErrors = append(processingErrors, err)
		}
	}

	logging.LogEvent("NotifyNewPosts", "status", "complete")

	return errors.Join(processingErrors...)
}
