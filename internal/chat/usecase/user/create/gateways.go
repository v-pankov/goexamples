package login

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Gateways interface {
	CreateNewUserEntity(ctx context.Context, userName user.Name) (*user.Entity, error)
	CreateNewUserEvent(ctx context.Context, userEntity *user.Entity) error
}
