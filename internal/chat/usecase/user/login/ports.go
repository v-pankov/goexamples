package login

import (
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type Request struct {
	UserName user.Name
}

type Response struct {
	SessionID session.ID
}
