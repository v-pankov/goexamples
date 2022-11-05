package msgbus

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type NewMessagesSessionSubscriber interface {
	SubscribeSessionForNewMessages(
		ctx context.Context,
		sessionID session.ID,
	) (
		<-chan *message.Entity,
		error,
	)
}