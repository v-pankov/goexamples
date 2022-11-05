package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type SessionFinder interface {
	FindSession(ctx context.Context, sessionID session.ID) (*session.Entity, error)
}
