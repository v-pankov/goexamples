package send

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
)

type Gateways struct {
	MessageBroadcaster messageGateway.Broadcaster
	MessageCreator     messageGateway.Creator
}
