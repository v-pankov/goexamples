package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type ActiveSessionCreator interface {
	CreateActiveSession(ctx context.Context, userID user.ID) (*session.Entity, error)
}
