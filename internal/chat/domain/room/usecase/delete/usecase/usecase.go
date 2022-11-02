package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/delete"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *delete.Args,
	) (
		*delete.Result,
		error,
	)
}

func New(
	gatewayUnsubscribeSessionsFromRoomMessages GatewayUnsubscribeSessionsFromRoomMessages,
	gatewayDeleteRoom GatewayDeleteRoom,
	gatewayNotifySessionsAboutRemovedRoom GatewayNotifySessionsAboutRemovedRoom,
) UseCase {
	return useCaseRoomDelete{
		gatewayUnsubscribeSessionsFromRoomMessages: gatewayUnsubscribeSessionsFromRoomMessages,
		gatewayDeleteRoom:                          gatewayDeleteRoom,
		gatewayNotifySessionsAboutRemovedRoom:      gatewayNotifySessionsAboutRemovedRoom,
	}
}

type useCaseRoomDelete struct {
	gatewayUnsubscribeSessionsFromRoomMessages GatewayUnsubscribeSessionsFromRoomMessages
	gatewayDeleteRoom                          GatewayDeleteRoom
	gatewayNotifySessionsAboutRemovedRoom      GatewayNotifySessionsAboutRemovedRoom
}

func (uc useCaseRoomDelete) Do(
	ctx context.Context,
	args *delete.Args,
) (
	*delete.Result,
	error,
) {
	err := uc.
		gatewayUnsubscribeSessionsFromRoomMessages.
		Call(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("unsubsribe sessions from room messages: %w", err)
	}

	err = uc.
		gatewayDeleteRoom.
		Call(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("delete room: %w", err)
	}

	err = uc.
		gatewayNotifySessionsAboutRemovedRoom.
		Call(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about removed room: %w", err)
	}

	return nil, nil
}
