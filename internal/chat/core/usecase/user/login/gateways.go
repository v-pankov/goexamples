package login

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity/user"
)

type Gateways interface {
	CreateOrFindUser(
		ctx context.Context,
		userName user.Name,
	) (
		*user.Entity,
		error,
	)
}
