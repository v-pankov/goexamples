package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/exit"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *exit.Args,
	) (
		*exit.Result,
		error,
	)
}

func New(
	gatewaySessionRoomMessagesUnsubscriber GatewaySessionRoomMessagesUnsubscriber,
) UseCase {
	return useCase{
		gatewaySessionRoomMessagesUnsubscriber: gatewaySessionRoomMessagesUnsubscriber,
	}
}

type useCase struct {
	gatewaySessionRoomMessagesUnsubscriber GatewaySessionRoomMessagesUnsubscriber
}

func (uc useCase) Do(
	ctx context.Context,
	args *exit.Args,
) (
	*exit.Result,
	error,
) {
	err := uc.
		gatewaySessionRoomMessagesUnsubscriber.
		GatewayUnsubscribeSessionForRoomMessages(
			ctx, args.SessionID, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("unsubscribe session for room messages: %w", err)
	}

	return &exit.Result{}, nil
}
