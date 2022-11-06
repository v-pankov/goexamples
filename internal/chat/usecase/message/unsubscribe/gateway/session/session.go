package session

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Finder interface {
	Find(ctx context.Context, sessionID session.ID) (*session.Entity, error)
}
