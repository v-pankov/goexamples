package session

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Creator interface {
	Create(ctx context.Context, userID user.ID) (*session.Entity, error)
}
