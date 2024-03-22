package repositories

import (
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"os"
	"strings"
)

type SubscriptionRepository interface {
	GetSubscriptions() ([]*models.Subscription, error)
}

func NewSubscriptionRepository() SubscriptionRepository {
	return &ConfigSubscriptionRepository{}
}

type ConfigSubscriptionRepository struct {
}

func (r *ConfigSubscriptionRepository) GetSubscriptions() ([]*models.Subscription, error) {
	blogs := strings.Split(os.Getenv("blog_monitor_targets"), ",")
	email := os.Getenv("blog_monitor_email")

	result := make([]*models.Subscription, 0)
	for _, item := range blogs {
		subscription := &models.Subscription{
			Id:                  "1",
			Name:                "Blog Post",
			BlogUrl:             item,
			NotificationMethod:  "email",
			NotificationTargets: []string{email},
			MaximumLookback:     1,
		}

		result = append(result, subscription)
	}

	return result, nil
}
