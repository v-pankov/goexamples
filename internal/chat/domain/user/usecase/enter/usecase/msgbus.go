package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter/usecase/msgbus"

type MessageBus interface {
	msgbus.NewMessagesSessionSubscriber
}
