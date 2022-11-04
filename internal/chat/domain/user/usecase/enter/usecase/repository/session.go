package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type SessionCreator interface {
	CreateSession(ctx context.Context, userID user.ID) (*session.Entity, error)
}
