package event

import (
	"context"
	"time"
)

type NewMessage struct {
	basicEvent
	MessageID       MessageID
	MessageContents MessageContents
}

type basicEvent struct {
	UserID    UserID
	Timestamp time.Time
}

type (
	UserID          string
	MessageID       int64
	MessageContents []byte
)

type Subscriber[Event any] interface {
	Subscribe(
		ctx context.Context,
		subscriptionID SubscriptionID,
	) (
		<-chan *Event,
		error,
	)
}

type Unsubscriber[Event any] interface {
	Unsubscribe(
		ctx context.Context,
		subscriptionID SubscriptionID,
	) error
}

type SubscriptionID string

type SubscriptionCreator interface {
	CreateSubscription(ctx context.Context) (SubscriptionID, error)
}

type SubscriptionDeleter interface {
	DeleteSubscription(ctx context.Context) (SubscriptionID, error)
}
