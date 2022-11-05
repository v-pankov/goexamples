package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/exit/usecase/msgbus"

type MessageBus interface {
	msgbus.NewMessagesSessionUnsubscriber
}
