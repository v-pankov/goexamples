package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
)

type GatewayUnsubscribeSessionsFromRoomMessages interface {
	Call(
		ctx context.Context, roomID room.ID,
	) error
}

type GatewayDeleteRoom interface {
	Call(
		ctx context.Context, roomID room.ID,
	) error
}

type GatewayNotifySessionsAboutRemovedRoom interface {
	Call(
		ctx context.Context, roomID room.ID,
	) error
}
