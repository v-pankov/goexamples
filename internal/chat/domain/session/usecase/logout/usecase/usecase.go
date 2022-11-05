package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout"
)

type UseCase interface {
	Do(ctx context.Context, args *logout.Args) (*logout.Result, error)
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

func (uc useCase) Do(ctx context.Context, args *logout.Args) (*logout.Result, error) {
	if err := uc.msgbus.UnsubscribeSessionFromNewMessages(ctx, args.SessionID); err != nil {
		return nil, fmt.Errorf("unsubsribe session from new messages: %w", err)
	}

	if err := uc.repository.DeactivateSession(ctx, args.SessionID); err != nil {
		return nil, fmt.Errorf("deactivate session: %w", err)
	}

	return &logout.Result{}, nil
}
