package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayFindUser interface {
	Call(
		ctx context.Context,
		userID user.ID,
	) (
		*user.Entity,
		error,
	)
}

type GatewayFindRoom interface {
	Call(
		ctx context.Context,
		roomName room.Name,
	) (
		*room.Entity,
		error,
	)
}
