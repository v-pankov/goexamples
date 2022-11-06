package unsubscribe

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/unsubscribe/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/unsubscribe/gateway/session"
)

type Gateways struct {
	MessageUnsubscriber messageGateway.Unsubscriber
	SessionFinder       sessionGateway.Finder
}
