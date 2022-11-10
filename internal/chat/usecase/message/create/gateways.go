package create

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Gateways interface {
	CreateNewMessageEntity(ctx context.Context, sessionID session.ID, messageText string) (*message.Entity, error)
	CreateNewMessageEvent(ctx context.Context, userID user.ID, message *message.Entity) error
}
