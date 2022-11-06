package message

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Subscriber interface {
	Subscribe(ctx context.Context, sessionID session.ID) (<-chan *message.Entity, error)
}
