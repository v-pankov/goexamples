package inmem

import (
	"context"
	"sync"
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity"

	usecaseMessageCreateGateways "github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/gateways"
)

type InMem struct {
	sync.Mutex
	messages      map[entity.MessageID]*entity.Message
	nextMessageID int64
}

func New() *InMem {
	return &InMem{
		messages: make(map[entity.MessageID]*entity.Message),
	}
}

func (inmem *InMem) UsecaseMessageCreateRepository() usecaseMessageCreateGateways.Repository {
	return uescaseMessageCreateRepositoryAdapter{inmem}
}

type uescaseMessageCreateRepositoryAdapter struct {
	*InMem
}

func (a uescaseMessageCreateRepositoryAdapter) CreateMessage(
	ctx context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	a.Lock()
	msg, err := a.createMessage(ctx, messageContents)
	a.Unlock()
	return msg, err
}

func (a uescaseMessageCreateRepositoryAdapter) createMessage(
	_ context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	a.nextMessageID++
	var (
		messageID = entity.MessageID(a.nextMessageID)
		message   = &entity.Message{
			ID:        messageID,
			Contents:  messageContents,
			CreatedAt: time.Now(),
		}
	)
	a.messages[messageID] = message
	return message, nil
}
