package message

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Creator interface {
	Create(ctx context.Context, sessionID session.ID, messageText string) (*message.Entity, error)
}

type Broadcaster interface {
	Broadcast(ctx context.Context, message *message.Entity) error
}
