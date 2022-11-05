package msgbus

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
)

type AllSessionsMessageBroadcaster interface {
	BroadcastMessageToAllSessions(ctx context.Context, message *message.Entity) error
}
