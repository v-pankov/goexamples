package message

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Unsubscriber interface {
	Unsubscribe(ctx context.Context, sessionID session.ID) error
}
