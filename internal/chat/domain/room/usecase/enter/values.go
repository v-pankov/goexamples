package enter

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type Args struct {
	SessionID session.ID
	RoomID    room.ID
}

type Result struct {
	Messages <-chan *message.Entity
}
