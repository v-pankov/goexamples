package send

import (
	"context"
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
	"github.com/vdrpkv/goexamples/internal/chat/infrastructure/repository/inmem"

	usecaseMessageSendGateways "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateways"
)

type Adapter struct {
	*inmem.InMem
}

var _ usecaseMessageSendGateways.Repository = Adapter{}

func (a Adapter) CreateMessage(
	ctx context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	a.Mutex.Lock()
	msg, err := a.createMessage(ctx, messageContents)
	a.Mutex.Unlock()
	return msg, err
}

func (a Adapter) createMessage(
	_ context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	var (
		messageID = entity.MessageID(a.NextMessageID)
		message   = &entity.Message{
			ID:        messageID,
			Contents:  messageContents,
			CreatedAt: time.Now(),
		}
	)
	a.Messages[messageID] = message
	a.NextMessageID++
	return message, nil
}
