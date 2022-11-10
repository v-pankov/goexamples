package login

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateways"
)

type Gateways interface {
	gateways.UserCreator
	gateways.UserFinder
	CreateSession(ctx context.Context, userID user.ID) (*session.Entity, error)
}
