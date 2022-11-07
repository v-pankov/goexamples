package unsubscribe

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/unsubscribe/gateway/message"
)

type Gateways struct {
	MessageUnsubscriber messageGateway.Unsubscriber
}
