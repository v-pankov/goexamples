package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
)

type GatewayRoomFinder interface {
	GatewayFindRoom(
		ctx context.Context, roomID room.ID,
	) (
		*room.Entity,
		error,
	)
}
