package create

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Gateways interface {
	CreateNewSessionEntity(ctx context.Context, userID user.ID) (*session.Entity, error)
	CreateNewSessionEvent(ctx context.Context, sessionEntity *session.Entity) error
}
