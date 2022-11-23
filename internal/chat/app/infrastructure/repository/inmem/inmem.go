package inmem

import (
	"sync"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity"
)

type InMem struct {
	Mutex         sync.Mutex
	Messages      map[entity.MessageID]*entity.Message
	NextMessageID int64
}

func New() *InMem {
	return &InMem{
		Messages:      make(map[entity.MessageID]*entity.Message),
		NextMessageID: 1,
	}
}
