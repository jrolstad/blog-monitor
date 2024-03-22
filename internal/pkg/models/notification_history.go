package models

import "time"

type NotificationHistory struct {
	SubscriptionId string
	PostId         string
	NotifiedAt     time.Time
}
