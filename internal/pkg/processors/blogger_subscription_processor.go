package processors

import (
	"context"
	"errors"
	"fmt"
	"github.com/jrolstad/blog-monitor/internal/pkg/logging"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
	"google.golang.org/api/blogger/v3"
	"google.golang.org/api/option"
	"time"
)

type BloggerSubscriptionProcessor struct {
	secretService services.SecretService
}

func (s *BloggerSubscriptionProcessor) ProcessSubscription(notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService,
	subscription *models.Subscription) error {
	logging.LogEvent("BloggerSubscriptionProcessor.ProcessSubscription", "subscription", subscription.Id)

	service, err := getGoogleService(s.secretService)
	if err != nil {
		return err
	}

	blogInfoRequest := service.Blogs.GetByUrl(subscription.BlogUrl)

	response, err := blogInfoRequest.Do()
	if err != nil {
		return err
	}
	logging.LogDependency("BloggerService", "action", "GetByUrl", "success", err == nil)

	posts, err := getLatestBlogPosts(service, response.Id, subscription.MaximumLookback)

	logging.LogEvent("GetSubscriptionPosts", "posts", len(posts))

	processingErrors := make([]error, 0)
	for _, item := range posts {
		alreadyNotified, err := notificationHistoryRepository.Exists(subscription.Id, item.Id)
		if err != nil {
			processingErrors = append(processingErrors, err)
			continue
		}
		logging.LogEvent("ProcessSubscriptionPost", "subscription", subscription.Id, "post", item.Id, "alreadyNotified", alreadyNotified)

		if !alreadyNotified {
			title := fmt.Sprintf("[%s] : %s", subscription.Name, item.Title)
			content := fmt.Sprintf("<p>Source: %s<p> %s", item.Url, item.Content)
			err = notificationService.Notify(subscription.NotificationMethod, subscription.NotificationTargets, title, content)
			if err != nil {
				processingErrors = append(processingErrors, err)
				continue
			}
			logging.LogEvent("ProcessSubscriptionNotified", "subscription", subscription.Id, "post", item.Id)

			err = notificationHistoryRepository.TrackNotification(subscription.Id, item.Id, time.Now())
			if err != nil {
				processingErrors = append(processingErrors, err)
			}
			logging.LogEvent("ProcessSubscriptionNotificationTracked", "subscription", subscription.Id, "post", item.Id)

		}
	}

	return errors.Join(processingErrors...)
}

func getGoogleService(secretService services.SecretService) (*blogger.Service, error) {
	ctx := context.Background()
	apiKey, err := secretService.GetGoogleApiKey()
	if err != nil {
		return nil, err
	}

	result, err := blogger.NewService(ctx, option.WithAPIKey(apiKey))
	logging.LogDependency("BloggerService", "action", "create", "success", err == nil)

	return result, err
}

func getLatestBlogPosts(service *blogger.Service, blogId string, maxPosts int) ([]*blogger.Post, error) {
	listRequest := service.Posts.List(blogId).
		FetchBodies(true).
		FetchImages(true).
		Status("live").
		View("READER").
		SortOption("descending").
		MaxResults(int64(maxPosts)).
		OrderBy("published")

	response, err := listRequest.Do()
	logging.LogDependency("BloggerService", "action", "GetPosts", "success", err == nil, "blogId", blogId)

	return response.Items, err
}
