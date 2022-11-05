package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *enter.Args,
	) (
		*enter.Result,
		error,
	)
}

func New(
	msgbus MessageBus,
	repository Repository,
) UseCase {
	return useCase{
		msgbus:     msgbus,
		repository: repository,
	}
}

type useCase struct {
	msgbus     MessageBus
	repository Repository
}

func (uc useCase) Do(
	ctx context.Context,
	args *enter.Args,
) (
	*enter.Result,
	error,
) {
	userEntity, err := uc.repository.CreateOrFindUser(ctx, args.UserName)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", args.UserName, err)
	}

	sessionEntity, err := uc.repository.CreateSession(ctx, userEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	messages, err := uc.msgbus.SubscribeForNewMessages(ctx, sessionEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("subscribe for new messages: %w", err)
	}

	return &enter.Result{
		Messages:  messages,
		SessionID: sessionEntity.ID,
	}, nil
}
