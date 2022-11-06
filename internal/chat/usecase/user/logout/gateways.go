package logout

import (
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/gateway/session"
)

type Gateways struct {
	SessionDeactivator sessionGateway.Deactivator
	SessionFinder      sessionGateway.Finder
}
