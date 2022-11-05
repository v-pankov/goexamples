package usecase

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/usecase/msgbus"
)

type MessageBus interface {
	msgbus.AllSessionsMessageBroadcaster
}
