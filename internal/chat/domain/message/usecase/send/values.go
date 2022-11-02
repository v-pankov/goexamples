package send

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type Args struct {
	AuthorUserID user.ID
	RoomID       room.ID
	MessageText  string
}

type Result struct {
}
