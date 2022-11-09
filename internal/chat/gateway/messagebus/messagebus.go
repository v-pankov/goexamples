package messagebus

import (
	usecaseMessageSubscribeGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/event/message/subscribe/gateway/message"
	usecaseMessageUnsubscribeGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/event/message/unsubscribe/gateway/message"
	usecaseMessageSendGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
)

type Gateways struct {
	UseCaseMessageSendGateways        UseCaseMessageSendGateways
	UseCaseMessageSubscribeGateways   UseCaseMessageSubscribeGateways
	UseCaseMessageUnsubscribeGateways UseCaseMessageUnsubscribeGateways
}

type UseCaseMessageSendGateways struct {
	Broadcaster usecaseMessageSendGateway.Broadcaster
}

type UseCaseMessageSubscribeGateways struct {
	Subscriber usecaseMessageSubscribeGateway.Subscriber
}

type UseCaseMessageUnsubscribeGateways struct {
	Unsubscriber usecaseMessageUnsubscribeGateway.Unsubscriber
}
