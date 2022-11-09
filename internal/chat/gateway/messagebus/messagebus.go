package messagebus

import (
	usecaseMessageSendGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
)

type Gateways struct {
	UseCaseMessageSendGateways UseCaseMessageSendGateways
}

type UseCaseMessageSendGateways struct {
	EventCreator usecaseMessageSendGateway.EventCreator
}
