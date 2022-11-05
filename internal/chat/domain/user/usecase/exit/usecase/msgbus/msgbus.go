package msgbus

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type NewMessagesSessionUnsubscriber interface {
	UnsubscribeSessionFromNewMessages(ctx context.Context, sessionID session.ID) error
}
