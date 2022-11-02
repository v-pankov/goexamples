package list

import "github.com/vdrpkv/goexamples/internal/chat/domain/room"

type Args struct {
}

type Result struct {
	Rooms []room.Entity
}
