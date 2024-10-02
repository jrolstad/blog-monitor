package processors

import (
	"errors"
	"fmt"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
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
		client, err := clients.NewNationalWeatherServiceClient()
		if err != nil {
			return nil, err
		}
		return &NationalWeatherServiceSubscriptionProcessor{client: client}, nil
	}
	return nil, errors.New(fmt.Sprintf("unrecognized type: |%s|", subscription.Type))
}
