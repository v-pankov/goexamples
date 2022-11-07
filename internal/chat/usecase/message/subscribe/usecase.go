package subscribe

import (
	"context"
	"fmt"
)

type UseCase interface {
	Do(ctx context.Context, request *Request) error
}

func New(
	gateways Gateways,
	presenter Presenter,
) UseCase {
	return useCase{
		gateways:  gateways,
		presenter: presenter,
	}
}

type useCase struct {
	gateways  Gateways
	presenter Presenter
}

func (uc useCase) Do(ctx context.Context, request *Request) error {
	messages, err := uc.gateways.MessageSubscriber.Subscribe(ctx, request.SessionID)
	if err != nil {
		return fmt.Errorf("subscribe messages: %w", err)
	}

	if err := uc.presenter.Present(ctx, &Response{
		Messages: messages,
	}); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
