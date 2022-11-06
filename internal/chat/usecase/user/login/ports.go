package login

import (
	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Request struct {
	UserName user.Name
}

type Response struct {
	Messages  <-chan *message.Entity
	SessionID session.ID
}
