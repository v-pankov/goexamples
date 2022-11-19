package create

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
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
