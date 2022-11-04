package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type UserFinder interface {
	FindUser(ctx context.Context, userID user.ID) (*user.Entity, error)
}
