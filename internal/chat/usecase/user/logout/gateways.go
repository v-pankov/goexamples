package logout

import (
	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/gateway/session"
)

type Gateways struct {
	MessageUnsubscriber messageGateway.Unsubscriber
	SessionDeactivator  sessionGateway.Deactivator
	SessionFinder       sessionGateway.Finder
}
