package subscribe

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/event/message"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/event/message/subscribe/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/event/message/subscribe/model/response"
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
	domainEvt, err := uc.gateways.MessageSubscriber.Subscribe(
		ctx, session.ID(requestCtx.SessionID),
	)
	if err != nil {
		return nil, fmt.Errorf("subscribe messages: %w", err)
	}

	// TODO: move it somewhere
	convertEventFn := func(evt *message.Event) *response.Event {
		modelEvt := &response.Event{
			UserID:    response.UserID(evt.Header.UserID),
			SessionID: response.SessionID(evt.Header.SessionID),
			MessageID: response.MessageID(evt.Header.MessageID),
			Time:      evt.Time,
		}
		switch evt.Type {
		case message.Created:
			modelEvt.Payload.New = &response.EventPayloadNew{
				MessageText: response.MessageText(evt.Data.Created.MessageContents),
			}
		case message.Updated:
			modelEvt.Payload.Edit = &response.EventPayloadEdit{
				MessageText: response.MessageText(evt.Data.Updated.MessageContents),
			}
		case message.Deleted:
			modelEvt.Payload.Delete = &response.EventPayloadDelete{}
		}
		return modelEvt
	}

	modelMessages := make(chan *response.Event)
	go func() {
		for msg := range domainEvt {
			modelMessages <- convertEventFn(msg)
		}
	}()

	return &response.Model{
		Messages: modelMessages,
	}, nil
}
