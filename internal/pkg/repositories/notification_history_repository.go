package repositories

import (
	"time"
)

type NotificationHistoryRepository interface {
	Exists(subscriptionId string, postId string) (bool, error)
	TrackNotification(subscriptionId string, postId string, notifiedAt time.Time) error
}

func NewNotificationHistoryRepository() NotificationHistoryRepository {
	return &InMemoryNotificationHistoryRepository{}
}

type InMemoryNotificationHistoryRepository struct {
}

func (*InMemoryNotificationHistoryRepository) Exists(subscriptionId string, postId string) (bool, error) {
	return false, nil
}

func (*InMemoryNotificationHistoryRepository) TrackNotification(subscriptionId string, postId string, notifiedAt time.Time) error {
	return nil
}
