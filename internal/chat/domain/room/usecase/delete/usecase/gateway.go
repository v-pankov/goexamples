package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
)

type GatewaySessionsRoomMessagesUnsubscriber interface {
	GatewayUnsubscribeSessionsFromRoomMessages(
		ctx context.Context, roomID room.ID,
	) error
}

type GatewayRoomDeleter interface {
	GatewayDeleteRoom(
		ctx context.Context, roomID room.ID,
	) error
}

type GatewaySessionsRoomRemovalNotifier interface {
	GatewayNotifySessionsAboutRemovedRoom(
		ctx context.Context, roomID room.ID,
	) error
}
