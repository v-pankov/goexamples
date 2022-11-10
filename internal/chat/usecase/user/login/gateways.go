package login

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Gateways interface {
	CreateOrFindUser(
		ctx context.Context,
		userName user.Name,
	) (
		*user.Entity,
		error,
	)

	CreateSession(
		ctx context.Context,
		userID user.ID,
	) (
		*session.Entity,
		error,
	)
}
