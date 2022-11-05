package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/login/usecase/msgbus"

type MessageBus interface {
	msgbus.NewMessagesSessionSubscriber
}
