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

func NewSubscriptionProcessor(subscription *models.Subscription, secretService services.SecretService) (SubscriptionProcessor, error) {
	if strings.EqualFold("blogger", subscription.Type) {
		return &BloggerSubscriptionProcessor{secretService: secretService}, nil
	}
	if strings.EqualFold("national-weather-service", subscription.Type) {
		return &NationalWeatherServiceSubscriptionProcessor{}, nil
	}
	return nil, errors.New(fmt.Sprintf("unrecognized type: |%s|", subscription.Type))
}
