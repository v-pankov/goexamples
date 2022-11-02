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
	gatewaySessionsRoomMessagesUnsubscriber GatewaySessionsRoomMessagesUnsubscriber,
	gatewayRoomDeleter GatewayRoomDeleter,
	gatewaySessionsRoomRemovalNotifier GatewaySessionsRoomRemovalNotifier,
) UseCase {
	return useCaseRoomDelete{
		gatewaySessionsRoomMessagesUnsubscriber: gatewaySessionsRoomMessagesUnsubscriber,
		gatewayRoomDeleter:                      gatewayRoomDeleter,
		gatewaySessionsRoomRemovalNotifier:      gatewaySessionsRoomRemovalNotifier,
	}
}

type useCaseRoomDelete struct {
	gatewaySessionsRoomMessagesUnsubscriber GatewaySessionsRoomMessagesUnsubscriber
	gatewayRoomDeleter                      GatewayRoomDeleter
	gatewaySessionsRoomRemovalNotifier      GatewaySessionsRoomRemovalNotifier
}

func (uc useCaseRoomDelete) Do(
	ctx context.Context,
	args *delete.Args,
) (
	*delete.Result,
	error,
) {
	err := uc.
		gatewaySessionsRoomMessagesUnsubscriber.
		GatewayUnsubscribeSessionsFromRoomMessages(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("unsubsribe sessions from room messages: %w", err)
	}

	err = uc.
		gatewayRoomDeleter.
		GatewayDeleteRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("delete room: %w", err)
	}

	err = uc.
		gatewaySessionsRoomRemovalNotifier.
		GatewayNotifySessionsAboutRemovedRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about removed room: %w", err)
	}

	return nil, nil
}
