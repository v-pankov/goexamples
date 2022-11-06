package login

import (
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/session"
	userGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/user"
)

type Gateways struct {
	SessionCreator    sessionGateway.Creator
	UserCreatorFinder userGateway.CreatorFinder
}
