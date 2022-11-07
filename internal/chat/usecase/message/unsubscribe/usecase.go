package unsubscribe

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
	if err := uc.gateways.MessageUnsubscriber.Unsubscribe(ctx, request.SessionID); err != nil {
		return fmt.Errorf("unsubsribe messages: %w", err)
	}

	if err := uc.presenter.Present(ctx, &Response{}); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
