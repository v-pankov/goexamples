package send

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Gateways interface {
	SendMessage(
		ctx context.Context,
		userID user.ID,
		message *message.Entity,
	) error
}
