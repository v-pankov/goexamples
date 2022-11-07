package session

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Deactivator interface {
	Deactivate(ctx context.Context, sessionID session.ID) error
}
