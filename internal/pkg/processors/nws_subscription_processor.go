package processors

import (
	"fmt"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
)

type NationalWeatherServiceSubscriptionProcessor struct {
	client *clients.NationalWeatherServiceClient
}

func (s *NationalWeatherServiceSubscriptionProcessor) ProcessSubscription(notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService,
	subscription *models.Subscription) error {

	logging.LogEvent("NationalWeatherServiceSubscriptionProcessor.ProcessSubscription", "subscription", subscription.Id)

	posts, err := s.client.GetData(subscription.BlogUrl)
	logging.LogDependency("NationalWeatherServiceClient", "action", "GetData", "success", err == nil)
	if err != nil {
		return err
	}

	mostRecentItem, err := getMostRecentItem(posts)
	if err != nil {
		return err
	}
	logging.LogEvent("GetMostRecentItem", "IssuedAt", mostRecentItem.IssuanceTime.Format("2006-01-02 15:04:05 MST"))

	alreadyNotified, err := notificationHistoryRepository.Exists(subscription.Id, mostRecentItem.URL)
	if err != nil {
		return err
	}
	if alreadyNotified {
		return nil
	}

	item, err := s.client.GetItem(mostRecentItem.URL)
	logging.LogDependency("NationalWeatherServiceClient", "action", "GetItem", "success", err == nil)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("[%s] : %s", subscription.Name, item.ProductName)
	content := fmt.Sprintf("<p>Source: %s<p> %s", item.URL, item.ProductText)
	err = notificationService.Notify(subscription.NotificationMethod, subscription.NotificationTargets, title, content)
	if err != nil {
		return err
	}
	logging.LogEvent("ProcessSubscriptionNotified", "subscription", subscription.Id, "post", item.URL)

	return nil
}

func getMostRecentItem(items []clients.NWSProduct) (clients.NWSProduct, error) {
	if len(items) == 0 {
		return clients.NWSProduct{}, fmt.Errorf("no products available")
	}

	mostRecent := items[0]
	for _, product := range items[1:] {
		if product.IssuanceTime.After(mostRecent.IssuanceTime) {
			mostRecent = product
		}
	}

	return mostRecent, nil
}
