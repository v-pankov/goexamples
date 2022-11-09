package send

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
)

type Gateways struct {
	EventCreator   messageGateway.EventCreator
	MessageCreator messageGateway.Creator
}
