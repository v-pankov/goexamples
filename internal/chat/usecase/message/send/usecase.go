package send

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
	messageEntity, err := uc.gateways.MessageCreator.Create(
		ctx,
		request.AuthorUserSessionID,
		request.MessageText,
	)
	if err != nil {
		return fmt.Errorf("create message: %w", err)
	}

	if err = uc.gateways.MessageBroadcaster.Broadcast(ctx, messageEntity); err != nil {
		return fmt.Errorf("broadcast message: %w", err)
	}

	if err := uc.presenter.Present(ctx, &Response{}); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
