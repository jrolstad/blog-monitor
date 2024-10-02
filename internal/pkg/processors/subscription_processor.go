package processors

import (
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

type SubscriptionProcessor interface {
	ProcessSubscription(notificationHistoryRepository repositories.NotificationHistoryRepository,
		notificationService services.NotificationService,
		subscription *models.Subscription) error
}

func NewSubscriptionProcessor(subscription *models.Subscription, secretService services.SecretService) (*BloggerSubscriptionProcessor, error) {

	return &BloggerSubscriptionProcessor{secretService: secretService}, nil
}
