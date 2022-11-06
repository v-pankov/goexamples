package login

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/session"
	userGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/user"
)

type Gateways struct {
	MessageSubscriber messageGateway.Subscriber
	SessionCreator    sessionGateway.Creator
	UserCreatorFinder userGateway.CreatorFinder
}
