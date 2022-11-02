package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
)

type GatewayFindRoom interface {
	Call(
		ctx context.Context, roomID room.ID,
	) (
		*room.Entity,
		error,
	)
}
