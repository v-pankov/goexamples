package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayCreateMessage interface {
	Call(
		ctx context.Context,
		authorUserID user.ID,
		roomID room.ID,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}

type GatewayNotifySessionsAboutNewMessage interface {
	Call(
		ctx context.Context,
		nessage *message.Entity,
	) error
}
