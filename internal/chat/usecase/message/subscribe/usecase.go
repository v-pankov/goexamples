package subscribe

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/subscribe/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/subscribe/model/response"
)

type UseCase interface {
	Do(
		ctx context.Context,
		requestCtx *request.Context,
		requestModel *request.Model,
	) (
		*response.Model,
		error,
	)
}

func New(
	gateways Gateways,
) UseCase {
	return useCase{
		gateways: gateways,
	}
}

type useCase struct {
	gateways Gateways
}

func (uc useCase) Do(
	ctx context.Context,
	requestCtx *request.Context,
	requestModel *request.Model,
) (
	*response.Model,
	error,
) {
	messages, err := uc.gateways.MessageSubscriber.Subscribe(
		ctx, session.ID(requestCtx.SessionID),
	)
	if err != nil {
		return nil, fmt.Errorf("subscribe messages: %w", err)
	}

	modelMessages := make(chan *response.Message)
	go func() {
		for msg := range messages {
			modelMessages <- &response.Message{
				SessionID: response.SessionID(msg.SessionID),
				Text:      msg.Text,
				CreatedAt: msg.CreatedAt,
			}
		}
	}()

	return &response.Model{
		Messages: modelMessages,
	}, nil
}
