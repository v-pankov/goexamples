package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayCreateRoom interface {
	Call(
		ctx context.Context,
		creatorUserID user.ID,
		roomName room.Name,
	) (
		*room.Entity,
		error,
	)
}

type GatewayNotifySessionsAboutNewRoom interface {
	Call(
		ctx context.Context,
		room *room.Entity,
	) error
}
