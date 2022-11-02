package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayCreateOrFindUser interface {
	Call(ctx context.Context, userName user.Name) (*user.Entity, error)
}

type GatewayCreateSession interface {
	Call(ctx context.Context, userID user.ID) (*session.Entity, error)
}
