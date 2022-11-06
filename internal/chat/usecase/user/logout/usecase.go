package logout

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
	sessionEntity, err := uc.gateways.SessionFinder.Find(ctx, request.SessionID)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrSessionNotFound
	}

	if err := uc.gateways.MessageUnsubscriber.Unsubscribe(ctx, sessionEntity.ID); err != nil {
		return fmt.Errorf("unsubsribe messages: %w", err)
	}

	if err := uc.gateways.SessionDeactivator.Deactivate(ctx, sessionEntity.ID); err != nil {
		return fmt.Errorf("deactivate session: %w", err)
	}

	if err := uc.presenter.Present(ctx, &Response{}); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
