package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayFindUser interface {
	Call(ctx context.Context, userID user.ID) (*user.Entity, error)
}
