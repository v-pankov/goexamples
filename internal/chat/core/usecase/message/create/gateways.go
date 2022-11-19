package create

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity/message"
)

type Gateways interface {
	CreateMessage(
		ctx context.Context,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}
