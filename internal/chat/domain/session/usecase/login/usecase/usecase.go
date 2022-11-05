package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/login"
)

type UseCase interface {
	Do(ctx context.Context, args *login.Args) (*login.Result, error)
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

func (uc useCase) Do(ctx context.Context, args *login.Args) (*login.Result, error) {
	userEntity, err := uc.repository.CreateOrFindUser(ctx, args.UserName)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", args.UserName, err)
	}

	sessionEntity, err := uc.repository.CreateActiveSession(ctx, userEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("create active session: %w", err)
	}

	messages, err := uc.msgbus.SubscribeSessionForNewMessages(ctx, sessionEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("subscribe for new messages: %w", err)
	}

	return &login.Result{
		Messages:  messages,
		SessionID: sessionEntity.ID,
	}, nil
}
