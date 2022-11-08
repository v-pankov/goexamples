package subscribe

import (
	"context"
	"fmt"
)

type UseCase interface {
	Do(ctx context.Context, request *Request) (*Response, error)
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

func (uc useCase) Do(ctx context.Context, request *Request) (*Response, error) {
	messages, err := uc.gateways.MessageSubscriber.Subscribe(ctx, request.SessionID)
	if err != nil {
		return nil, fmt.Errorf("subscribe messages: %w", err)
	}

	return &Response{
		Messages: messages,
	}, nil
}
