package subscribe

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/subscribe/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/subscribe/gateway/session"
)

type Gateways struct {
	MessageSubscriber messageGateway.Subscriber
	SessionFinder     sessionGateway.Finder
}
