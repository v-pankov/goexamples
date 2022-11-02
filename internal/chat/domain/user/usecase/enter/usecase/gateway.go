package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type GatewayUserCreatorFinder interface {
	GatewayCreateOrFindUser(ctx context.Context, userName user.Name) (*user.Entity, error)
}

type GatewaySessionCreator interface {
	GatewayCreateSession(ctx context.Context, userID user.ID) (*session.Entity, error)
}
