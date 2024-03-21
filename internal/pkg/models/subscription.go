package models

type Subscription struct {
	BlogUrl             string
	NotificationMethod  string
	NotificationTargets []string
	MaximumLookback     int
}
