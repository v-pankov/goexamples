package usecase

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/usecase/messagebus"
)

type MessageBus interface {
	messagebus.RoomMessageBroadcaster
}
