package logout

import "github.com/vdrpkv/goexamples/internal/chat/domain/session"

type Args struct {
	SessionID session.ID
}

type Result struct {
}
