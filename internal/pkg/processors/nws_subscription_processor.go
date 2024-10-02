package processors

import (
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

type NationalWeatherServiceSubscriptionProcessor struct {
}

func (s *NationalWeatherServiceSubscriptionProcessor) ProcessSubscription(notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService,
	subscription *models.Subscription) error {

	logging.LogEvent("NationalWeatherServiceSubscriptionProcessor.ProcessSubscription", "subscription", subscription.Id)

	return nil
}
