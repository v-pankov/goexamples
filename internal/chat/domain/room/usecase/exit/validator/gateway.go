package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type GatewaySessionFinder interface {
	GatewayFindSession(
		ctx context.Context, sessionID session.ID,
	) (
		*session.Entity,
		error,
	)
}

type GatewayRoomFinder interface {
	GatewayFindRoom(
		ctx context.Context, roomID room.ID,
	) (
		*room.Entity,
		error,
	)
}
