package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayUserFinder interface {
	GatewayFindUser(ctx context.Context, userID user.ID) (*user.Entity, error)
}

type GatewayRoomFinder interface {
	GatewayRoomFinder(ctx context.Context, roomID room.ID) (*room.Entity, error)
}
