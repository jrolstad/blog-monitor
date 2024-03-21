package repositories

import (
	"time"
)

type NotificationHistoryRepository interface {
	TrackNotification(subscriptionId string, postId string, notifiedAt time.Time) error
}

func NewNotificationHistoryRepository() NotificationHistoryRepository {
	return &InMemoryNotificationHistoryRepository{}
}

type InMemoryNotificationHistoryRepository struct {
}

func (*InMemoryNotificationHistoryRepository) TrackNotification(subscriptionId string, postId string, notifiedAt time.Time) error {
	return nil
}
