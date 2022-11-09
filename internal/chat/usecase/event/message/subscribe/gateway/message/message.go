package message

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/event/message"
)

type Subscriber interface {
	Subscribe(ctx context.Context, sessionID session.ID) (<-chan *message.Event, error)
}
