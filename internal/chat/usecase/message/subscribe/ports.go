package subscribe

import (
	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type Request struct {
	SessionID session.ID
}

type Response struct {
	Messages <-chan *message.Entity
}
