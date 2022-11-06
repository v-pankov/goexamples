package login

import (
	"context"
	"fmt"
	"strings"
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
	if len(strings.TrimSpace(request.UserName.String())) == 0 {
		return ErrEmptyUserName
	}

	userEntity, err := uc.gateways.UserCreatorFinder.CreateOrFind(ctx, request.UserName)
	if err != nil {
		return fmt.Errorf("create or find user [%s]: %w", request.UserName, err)
	}

	sessionEntity, err := uc.gateways.SessionCreator.Create(ctx, userEntity.ID)
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}

	messages, err := uc.gateways.MessageSubscriber.Subscribe(ctx, sessionEntity.ID)
	if err != nil {
		return fmt.Errorf("subscribe messages: %w", err)
	}

	if err := uc.presenter.Present(ctx, &Response{
		Messages:  messages,
		SessionID: sessionEntity.ID,
	}); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
