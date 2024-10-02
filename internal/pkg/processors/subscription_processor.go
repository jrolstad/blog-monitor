package processors

import (
	"errors"
	"fmt"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
	"strings"
)

type SubscriptionProcessor interface {
	ProcessSubscription(notificationHistoryRepository repositories.NotificationHistoryRepository,
		notificationService services.NotificationService,
		subscription *models.Subscription) error
}

func NewSubscriptionProcessor(subscription *models.Subscription, secretService services.SecretService) (*BloggerSubscriptionProcessor, error) {
	if strings.EqualFold("blogger", subscription.Type) {
		return &BloggerSubscriptionProcessor{secretService: secretService}, nil
	}
	return nil, errors.New(fmt.Sprintf("Unrecognized type: |%s|", subscription.Type))
}
