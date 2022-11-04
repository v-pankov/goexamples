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
	gatewaySubscribeSessionForRoomMessages GatewaySubscribeSessionForRoomMessages,
) UseCase {
	return useCase{
		gatewaySubscribeSessionForRoomMessages: gatewaySubscribeSessionForRoomMessages,
	}
}

type useCase struct {
	gatewaySubscribeSessionForRoomMessages GatewaySubscribeSessionForRoomMessages
}

func (uc useCase) Do(
	ctx context.Context,
	args *enter.Args,
) (
	*enter.Result,
	error,
) {
	messages, err := uc.
		gatewaySubscribeSessionForRoomMessages.
		Call(
			ctx, args.SessionID, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("subscribe session for room messages: %w", err)
	}

	return &enter.Result{
		Messages: messages,
	}, nil
}
