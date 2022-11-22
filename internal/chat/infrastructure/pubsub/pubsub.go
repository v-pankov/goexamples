package pubsub

import "github.com/vdrpkv/goexamples/internal/chat/core"

type (
	Sub interface {
		core.Receiver
		Unsubscribe() error
	}
)

type Pub interface {
	core.Sender
	Subscribe() (Sub, error)
}
