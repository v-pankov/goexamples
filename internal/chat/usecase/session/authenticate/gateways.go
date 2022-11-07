package authenticate

import (
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/session/authenticate/gateway/session"
)

type Gateways struct {
	SessionFinder sessionGateway.Finder
}
