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
	gatewayMessageCreator GatewayMessageCreator,
	gatewayNewMessageSessionsNotifier GatewayNewMessageSessionsNotifier,
) UseCase {
	return useCase{
		gatewayMessageCreator:             gatewayMessageCreator,
		gatewayNewMessageSessionsNotifier: gatewayNewMessageSessionsNotifier,
	}
}

type useCase struct {
	gatewayMessageCreator             GatewayMessageCreator
	gatewayNewMessageSessionsNotifier GatewayNewMessageSessionsNotifier
}

func (uc useCase) Do(
	ctx context.Context,
	args *send.Args,
) (
	*send.Result,
	error,
) {
	messageEntity, err := uc.
		gatewayMessageCreator.
		GatewayCreateMessage(
			ctx,
			args.AuthorUserID,
			args.RoomID,
			args.MessageText,
		)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	err = uc.
		gatewayNewMessageSessionsNotifier.
		GatewayNotifySessionsAboutNewMessage(
			ctx, messageEntity,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about new message: %w", err)
	}

	return &send.Result{}, nil
}
