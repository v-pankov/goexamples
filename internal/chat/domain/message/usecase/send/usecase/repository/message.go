package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type MessageCreator interface {
	CreateMessage(
		ctx context.Context,
		authorUserSessionID session.ID,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}
