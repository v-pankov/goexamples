package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayMessageCreator interface {
	GatewayCreateMessage(
		ctx context.Context,
		authorUserID user.ID,
		roomID room.ID,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}

type GatewayNewMessageSessionsNotifier interface {
	GatewayNotifySessionsAboutNewMessage(
		ctx context.Context,
		nessage *message.Entity,
	) error
}
