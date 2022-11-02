package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type GatewayFindSession interface {
	Call(
		ctx context.Context, sessionID session.ID,
	) (
		*session.Entity,
		error,
	)
}
