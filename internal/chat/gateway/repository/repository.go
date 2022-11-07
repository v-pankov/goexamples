package repository

import (
	usecaseMessageSendMessageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"

	usecaseSessionAuthenticateSessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/session/authenticate/gateway/session"

	usecaseUserLoginSessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/session"
	usecaseUserLoginUserGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/user"

	usecaseUserLogoutSessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/gateway/session"
)

type Gateways struct {
	UseCaseMessageSendGateways         UseCaseMessageSendGateways
	UseCaseSessionAuthenticateGateways UseCaseSessionAuthenticateGateways
	UseCaseUserLoginGateways           UseCaseUserLoginGateways
	UseCaseUserLogoutGateways          UseCaseUserLogoutGateways
}

type UseCaseMessageSendGateways struct {
	MessageCreator usecaseMessageSendMessageGateway.Creator
}

type UseCaseSessionAuthenticateGateways struct {
	SessionFinder usecaseSessionAuthenticateSessionGateway.Finder
}

type UseCaseUserLoginGateways struct {
	SessionCreator    usecaseUserLoginSessionGateway.Creator
	UserCreatorFinder usecaseUserLoginUserGateway.CreatorFinder
}

type UseCaseUserLogoutGateways struct {
	SessionDeactivator usecaseUserLogoutSessionGateway.Deactivator
}
