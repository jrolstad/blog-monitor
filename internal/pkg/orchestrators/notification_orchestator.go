package orchestrators

import (
	"context"
	"errors"
	"github.com/jrolstad/blog-monitor/internal/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"github.com/jrolstad/blog-monitor/internal/pkg/repositories"
	"github.com/jrolstad/blog-monitor/internal/pkg/services"
	"google.golang.org/api/blogger/v3"
	"google.golang.org/api/option"
	"time"
)

func NotifyNewPosts(config *configuration.AppConfig,
	subscriptionRepository repositories.SubscriptionRepository,
	notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService) error {
	service, err := getGoogleService(config)
	if err != nil {
		return err
	}

	subscriptions, err := subscriptionRepository.GetSubscriptions()
	if err != nil {
		return err
	}

	processingErrors := make([]error, 0)
	for _, item := range subscriptions {
		err = processSubscription(service, notificationHistoryRepository, notificationService, item)
		if err != nil {
			processingErrors = append(processingErrors, err)
		}
	}

	return errors.Join(processingErrors...)
}

func getGoogleService(config *configuration.AppConfig) (*blogger.Service, error) {
	ctx := context.Background()
	return blogger.NewService(ctx, option.WithAPIKey(config.GoogleApiKey))
}

func processSubscription(service *blogger.Service,
	notificationHistoryRepository repositories.NotificationHistoryRepository,
	notificationService services.NotificationService,
	subscription *models.Subscription) error {
	blogInfoRequest := service.Blogs.GetByUrl(subscription.BlogUrl)

	response, err := blogInfoRequest.Do()
	if err != nil {
		return err
	}

	posts, err := getLatestBlogPosts(service, response.Id, subscription.MaximumLookback)

	processingErrors := make([]error, 0)
	for _, item := range posts {
		err = notificationService.Notify(subscription.NotificationMethod, subscription.NotificationTargets, item.Title, item.Content)
		if err != nil {
			processingErrors = append(processingErrors, err)
		}

		err = notificationHistoryRepository.TrackNotification(subscription.Id, item.Id, time.Now())
		if err != nil {
			processingErrors = append(processingErrors, err)
		}
	}

	return errors.Join(processingErrors...)
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
	return response.Items, err
}
