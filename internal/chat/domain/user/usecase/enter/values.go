package enter

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type Args struct {
	UserName user.Name
}

type Result struct {
	SessionID session.ID
}
