package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type GatewaySessionRoomMessagesSubscriber interface {
	GatewaySubscribeSessionForRoomMessages(
		ctx context.Context,
		sessionID session.ID,
		roomID room.ID,
	) error
}
