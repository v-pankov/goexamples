package repository

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type UserCreatorFinder interface {
	CreateOrFindUser(ctx context.Context, userName user.Name) (*user.Entity, error)
}
