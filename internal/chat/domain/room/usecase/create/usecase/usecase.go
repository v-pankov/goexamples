package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/create"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *create.Args,
	) (
		*create.Result,
		error,
	)
}

func New(
	gatewayRoomCreator GatewayRoomCreator,
	gatewayNewRoomSessionsNotifier GatewayNewRoomSessionsNotifier,
) UseCase {
	return useCase{
		gatewayRoomCreator:             gatewayRoomCreator,
		gatewayNewRoomSessionsNotifier: gatewayNewRoomSessionsNotifier,
	}
}

type useCase struct {
	gatewayRoomCreator             GatewayRoomCreator
	gatewayNewRoomSessionsNotifier GatewayNewRoomSessionsNotifier
}

func (uc useCase) Do(
	ctx context.Context,
	args *create.Args,
) (
	*create.Result,
	error,
) {
	roomEntity, err := uc.
		gatewayRoomCreator.
		GatewayCreateRoom(
			ctx, args.CreatorUserID, args.RoomName,
		)
	if err != nil {
		return nil, fmt.Errorf("create room: %w", err)
	}

	err = uc.
		gatewayNewRoomSessionsNotifier.
		GatewayNotifySessionsAboutNewRoom(
			ctx, roomEntity,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about new room: %w", err)
	}

	return &create.Result{
		Room: roomEntity,
	}, nil
}
