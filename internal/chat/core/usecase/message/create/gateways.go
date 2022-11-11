package create

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/core/entity/user"
)

type Gateways interface {
	CreateMessage(
		ctx context.Context,
		userID user.ID,
		messageText string,
	) (
		*message.Entity,
		error,
	)
}
