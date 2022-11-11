package logout

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity/session"
)

type Gateways interface {
	DeleteSession(
		ctx context.Context,
		sessionID session.ID,
	) error
}
