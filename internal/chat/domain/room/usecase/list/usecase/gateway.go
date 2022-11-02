package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
)

type GatewayGetAllRooms interface {
	Call(
		ctx context.Context,
	) (
		[]room.Entity,
		error,
	)
}
