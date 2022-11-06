package send

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/session"
)

type Gateways struct {
	MessageBroadcaster messageGateway.Broadcaster
	MessageCreator     messageGateway.Creator
	SessionFinder      sessionGateway.Finder
}
