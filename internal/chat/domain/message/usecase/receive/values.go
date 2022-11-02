package receive

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type Args struct {
	SessionID session.ID
	Message   *message.Entity
}

type Result struct {
}
