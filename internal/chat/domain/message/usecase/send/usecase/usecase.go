package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send"
)

type UseCase interface {
	Do(ctx context.Context, args *send.Args) (*send.Result, error)
}

func New(
	repository Repository,
	messageBus MessageBus,
) UseCase {
	return useCase{
		repository: repository,
		messageBus: messageBus,
	}
}

type useCase struct {
	repository Repository
	messageBus MessageBus
}

func (uc useCase) Do(
	ctx context.Context,
	args *send.Args,
) (
	*send.Result,
	error,
) {

	messageEntity, err := uc.repository.CreateMessage(
		ctx,
		args.AuthorUserID,
		args.MessageText,
	)

	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	err = uc.messageBus.BroadcastMessageToAllSessions(ctx, messageEntity)

	if err != nil {
		return nil, fmt.Errorf("broadcast message to all sessions: %w", err)
	}

	return &send.Result{}, nil
}
