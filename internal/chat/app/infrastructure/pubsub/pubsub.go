package pubsub

import appIO "github.com/vdrpkv/goexamples/internal/chat/app/io"

type (
	Sub interface {
		appIO.Receiver
		Unsubscribe() error
	}
)

type Pub interface {
	appIO.Sender
	Subscribe() (Sub, error)
}
