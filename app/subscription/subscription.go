package subscription

import "github.com/v4fly/v4ray-core/v0/features"

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

type SubscriptionManager interface {
	features.Feature
	AddTrackedSubscriptionFromImportSource(importSource *ImportSource) error
	RemoveTrackedSubscription(name string) error
	ListTrackedSubscriptions() []string
	GetTrackedSubscriptionStatus(name string) (*TrackedSubscriptionStatus, error)
	UpdateTrackedSubscription(name string) error
}

func SubscriptionManagerType() interface{} {
	return (*SubscriptionManager)(nil)
}
