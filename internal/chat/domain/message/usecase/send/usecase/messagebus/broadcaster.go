package messagebus

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
)

type RoomMessageBroadcaster interface {
	BroadcastMessageInARoom(ctx context.Context, message *message.Entity) error
}
