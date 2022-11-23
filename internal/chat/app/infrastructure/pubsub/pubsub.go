package pubsub

import "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/transport"

type (
	Sub interface {
		transport.Receiver
		Unsubscribe() error
	}
)

type Pub interface {
	transport.Sender
	Subscribe() (Sub, error)
}
