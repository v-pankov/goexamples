package inmem

import (
	"context"
	"sync"
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/gateways"
)

type inmem struct {
	sync.Mutex
	messages      map[entity.MessageID]*entity.Message
	nextMessageID int64
}

func New() gateways.Repository {
	return &inmem{
		messages: make(map[entity.MessageID]*entity.Message),
	}
}

func (inmem *inmem) CreateMessage(
	ctx context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	inmem.Lock()
	msg, err := inmem.createMessage(ctx, messageContents)
	inmem.Unlock()
	return msg, err
}

func (inmem *inmem) createMessage(
	_ context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	inmem.nextMessageID++
	var (
		messageID = entity.MessageID(inmem.nextMessageID)
		message   = &entity.Message{
			ID:        messageID,
			Contents:  messageContents,
			CreatedAt: time.Now(),
		}
	)
	inmem.messages[messageID] = message
	return message, nil
}
