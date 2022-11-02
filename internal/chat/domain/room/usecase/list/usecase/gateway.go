package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
)

type GatewayRoomGetter interface {
	GatewayGetAllRooms(
		ctx context.Context,
	) (
		[]room.Entity,
		error,
	)
}
