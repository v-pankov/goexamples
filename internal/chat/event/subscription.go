package event

import "context"

type SubscriptionID int64

type SubscriptionCreator interface {
	CreateSubscription(ctx context.Context) (SubscriptionID, error)
}

type SubscriptionDeleter interface {
	DeleteSubscription(ctx context.Context, subscriptionID SubscriptionID) error
}

type Subscriber[Event any] interface {
	Subscribe(ctx context.Context, subscriptionID SubscriptionID) (<-chan *Event, error)
}

type Unsubscriber[Event any] interface {
	Unsubsribe(ctx context.Context, subscriptionID SubscriptionID) error
}

type Publisher[Event any] interface {
	Publish(ctx context.Context, evt *Event, subsriptionIDs ...SubscriptionID) error
}
