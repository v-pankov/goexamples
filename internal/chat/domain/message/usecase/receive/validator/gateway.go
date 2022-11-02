package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type GatewaySessionFinder interface {
	GatewayFindSession(
		ctx context.Context, sessionID session.ID,
	) (
		*session.Entity,
		error,
	)
}
