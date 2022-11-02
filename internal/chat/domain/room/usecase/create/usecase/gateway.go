package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayRoomCreator interface {
	GatewayCreateRoom(
		ctx context.Context,
		creatorUserID user.ID,
		roomName room.Name,
	) (
		*room.Entity,
		error,
	)
}

type GatewayNewRoomSessionsNotifier interface {
	GatewayNotifySessionsAboutNewRoom(
		ctx context.Context,
		room *room.Entity,
	) error
}
