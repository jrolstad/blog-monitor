package processors

import (
	"fmt"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
	"html"
	"regexp"
	"strings"
	"time"
)

type NationalWeatherServiceSubscriptionProcessor struct {
	client *clients.NationalWeatherServiceClient
}

func (s *NationalWeatherServiceSubscriptionProcessor) ProcessSubscription(notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService,
	subscription *models.Subscription) error {

	logging.LogEvent("NationalWeatherServiceSubscriptionProcessor.ProcessSubscription", "subscription", subscription.Id)

	posts, err := s.client.GetProductData(subscription.BlogUrl)
	logging.LogDependency("NationalWeatherServiceClient", "action", "GetData", "success", err == nil)
	if err != nil {
		return err
	}
	if len(posts) == 0 {
		logging.LogEvent("No Items Received")
		return nil
	}

	mostRecentItem, err := getMostRecentItem(posts)
	if err != nil {
		return err
	}
	logging.LogEvent("GetMostRecentItem", "IssuedAt", formatDate(mostRecentItem.IssuanceTime))

	alreadyNotified, err := notificationHistoryRepository.Exists(subscription.Id, mostRecentItem.URL)
	if err != nil {
		return err
	}
	if alreadyNotified {
		return nil
	}

	detailedItem, err := s.client.GetProductItem(mostRecentItem.URL)
	logging.LogDependency("NationalWeatherServiceClient", "action", "GetItem", "success", err == nil)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("[%s] : %s", subscription.Name, formatDate(detailedItem.IssuanceTime))
	formattedProductText := formatProductTextToHtml(detailedItem.ProductText)
	content := fmt.Sprintf("<p>Source: %s<p> %s", detailedItem.URL, formattedProductText)

	err = notificationService.Notify(subscription.NotificationMethod, subscription.NotificationTargets, title, content)
	if err != nil {
		return err
	}
	logging.LogEvent("ProcessSubscriptionNotified", "subscription", subscription.Id, "post", detailedItem.URL)

	err = notificationHistoryRepository.TrackNotification(subscription.Id, detailedItem.URL, time.Now())
	if err != nil {
		return err
	}
	logging.LogEvent("ProcessSubscriptionNotificationTracked", "subscription", subscription.Id, "post", detailedItem.URL)

	return nil
}

func formatDate(toFormat time.Time) string {
	return toFormat.Format(time.RFC850)
}

func getMostRecentItem(items []clients.NWSProduct) (clients.NWSProduct, error) {
	if len(items) == 0 {
		return clients.NWSProduct{}, nil
	}

	mostRecent := items[0]
	for _, product := range items[1:] {
		if product.IssuanceTime.After(mostRecent.IssuanceTime) {
			mostRecent = product
		}
	}

	return mostRecent, nil
}

func formatProductTextToHtml(text string) string {
	escapedText := html.EscapeString(text)

	sectionRegex := regexp.MustCompile(`(?m)^(\.\w[\w\s/]+)\.\.\.`)
	splitSections := sectionRegex.Split(escapedText, -1)
	sectionMatches := sectionRegex.FindAllStringSubmatch(escapedText, -1)

	var htmlBuilder strings.Builder
	for sectionIndex, section := range splitSections {
		if isSectionHeader(sectionIndex, sectionMatches) {
			header := strings.TrimPrefix(sectionMatches[sectionIndex-1][1], ".")
			formattedSectionHeader := fmt.Sprintf("<h2>%s</h2>", html.EscapeString(header))
			htmlBuilder.WriteString(formattedSectionHeader)
		}

		cleanedSection := removeArtificialNewlinesInSection(section)
		formattedSectionDetail := strings.ReplaceAll(cleanedSection, "\n", "<br/>")
		htmlBuilder.WriteString(formattedSectionDetail)
	}

	return htmlBuilder.String()
}

func isSectionHeader(sectionIndex int, sectionMatches [][]string) bool {
	return sectionIndex > 0 && sectionIndex <= len(sectionMatches)
}

func removeArtificialNewlinesInSection(section string) string {
	expression := regexp.MustCompile(`([^\n])\n([^\n])`)
	return expression.ReplaceAllString(section, "$1 $2")
}
