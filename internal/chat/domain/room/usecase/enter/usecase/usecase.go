package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/enter"
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
	gatewaySessionRoomMessagesSubscriber GatewaySessionRoomMessagesSubscriber,
) UseCase {
	return useCase{
		gatewaySessionRoomMessagesSubscriber: gatewaySessionRoomMessagesSubscriber,
	}
}

type useCase struct {
	gatewaySessionRoomMessagesSubscriber GatewaySessionRoomMessagesSubscriber
}

func (uc useCase) Do(
	ctx context.Context,
	args *enter.Args,
) (
	*enter.Result,
	error,
) {
	err := uc.
		gatewaySessionRoomMessagesSubscriber.
		GatewaySubscribeSessionForRoomMessages(
			ctx, args.SessionID, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("subscribe session for room messages: %w", err)
	}

	return &enter.Result{}, nil
}
