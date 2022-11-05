package login

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type Args struct {
	UserName user.Name
}

type Result struct {
	Messages  <-chan *message.Entity
	SessionID session.ID
}
