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
	authorUserSessionEntity, err := uc.gateways.SessionFinder.Find(ctx, request.AuthorUserSessionID)
	if err != nil {
		return fmt.Errorf("find author user session: %w", err)
	}

	if authorUserSessionEntity == nil {
		return ErrAuthorUserSessionNotFound
	}

	messageEntity, err := uc.gateways.MessageCreator.Create(
		ctx,
		request.AuthorUserSessionID,
		request.MessageText,
	)

	if err != nil {
		return fmt.Errorf("create message: %w", err)
	}

	err = uc.gateways.MessageBroadcaster.Broadcast(ctx, messageEntity)

	if err != nil {
		return fmt.Errorf("broadcast message to all sessions: %w", err)
	}

	if err := uc.presenter.Present(ctx, &Response{}); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
