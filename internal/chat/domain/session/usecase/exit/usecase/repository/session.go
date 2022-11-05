package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type SessionDeactivator interface {
	DeactivateSession(ctx context.Context, sessionID session.ID) error
}
