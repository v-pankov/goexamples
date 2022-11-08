package data

import (
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type New struct {
	UserID      user.ID
	SessionID   session.ID
	MessageText string
}

type Edit struct {
	UserID      user.ID
	SessionID   session.ID
	MessageText string
}

type Delete struct {
	UserID    user.ID
	SessionID session.ID
}
