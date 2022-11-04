package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type MessageCreator interface {
	CreateMessage(
		ctx context.Context,
		authorUserID user.ID,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}
