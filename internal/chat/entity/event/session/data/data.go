package data

import (
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type New struct {
	SessionID session.ID
	UserID    user.ID
}

type Delete struct {
	SessionID session.ID
	UserID    user.ID
}
