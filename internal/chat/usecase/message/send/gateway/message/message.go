package message

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Creator interface {
	Create(ctx context.Context, sessionID session.ID, messageText string) (*message.Entity, error)
}

type EventCreator interface {
	CreateEvent(ctx context.Context, userID user.ID, message *message.Entity) error
}
