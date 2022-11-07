package subscribe

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/subscribe/gateway/message"
)

type Gateways struct {
	MessageSubscriber messageGateway.Subscriber
}
