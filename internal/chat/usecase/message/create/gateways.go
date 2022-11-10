package create

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Gateways interface {
	CreateMessage(
		ctx context.Context,
		sessionID session.ID,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}
