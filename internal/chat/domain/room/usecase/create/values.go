package create

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

type Args struct {
	CreatorUserID user.ID
	RoomName      room.Name
}

type Result struct {
	Room *room.Entity
}
