package delete

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Gateways interface {
	FindSessionEntity(ctx context.Context, sessionID session.ID) (*session.Entity, error)
	UpdateSessionEntity(ctx context.Context, sessionEntity *session.Entity) error
	CreateSessionDeletedEvent(ctx context.Context, sessionEntity *session.Entity) error
}
