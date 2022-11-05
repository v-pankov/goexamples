package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/usecase/msgbus"

type MessageBus interface {
	msgbus.NewMessagesSessionUnsubscriber
}
