package user

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type CreatorFinder interface {
	CreateOrFind(ctx context.Context, userName user.Name) (*user.Entity, error)
}
