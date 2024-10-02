package models

type Subscription struct {
	Id                  string
	Name                string
	BlogUrl             string
	NotificationMethod  string
	NotificationTargets []string
	MaximumLookback     int
	Type                string
}
