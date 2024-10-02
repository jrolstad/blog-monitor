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

	detailedItem, err := s.client.GetItem(mostRecentItem.URL)
	logging.LogDependency("NationalWeatherServiceClient", "action", "GetItem", "success", err == nil)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("[%s] : %s", subscription.Name, detailedItem.ProductName)
	formattedProductText := formatProductTextToHTML(detailedItem.ProductText)
	content := fmt.Sprintf("<p>Source: %s<p> %s", detailedItem.URL, formattedProductText)

	err = notificationService.Notify(subscription.NotificationMethod, subscription.NotificationTargets, title, content)
	if err != nil {
		return err
	}
	logging.LogEvent("ProcessSubscriptionNotified", "subscription", subscription.Id, "post", detailedItem.URL)

	//err = notificationHistoryRepository.TrackNotification(subscription.Id, detailedItem.URL, time.Now())
	//if err != nil {
	//	return err
	//}
	logging.LogEvent("ProcessSubscriptionNotificationTracked", "subscription", subscription.Id, "post", detailedItem.URL)

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

func formatProductTextToHTML(text string) string {
	// Escape HTML special characters
	escapedText := html.EscapeString(text)

	// Define a regular expression to capture section headers (e.g., ".SYNOPSIS", ".SHORT TERM", etc.)
	sectionRegex := regexp.MustCompile(`(?m)^(\.\w[\w\s/]+)\.\.\.`)

	// Split text into sections and retain delimiters
	splitSections := sectionRegex.Split(escapedText, -1)
	sectionMatches := sectionRegex.FindAllStringSubmatch(escapedText, -1)

	// Start building the HTML content
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<html><body>")

	// Handle sections
	for i, section := range splitSections {
		if i > 0 && i <= len(sectionMatches) { // Add headers for each section
			header := strings.TrimPrefix(sectionMatches[i-1][1], ".")
			htmlBuilder.WriteString(fmt.Sprintf("<h2>%s</h2>", html.EscapeString(header)))
		}
		// Replace newlines with <br/> and add section text
		sectionHTML := strings.ReplaceAll(section, "\n", "<br/>")
		htmlBuilder.WriteString(sectionHTML)
	}

	htmlBuilder.WriteString("</body></html>")
	return htmlBuilder.String()
}
