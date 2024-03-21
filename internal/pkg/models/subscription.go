package models

type Subscription struct {
	Id                  string
	BlogUrl             string
	NotificationMethod  string
	NotificationTargets []string
	MaximumLookback     int
}
