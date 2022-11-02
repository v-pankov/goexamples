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
	gatewayCreateRoom GatewayCreateRoom,
	gatewayNotifySessionsAboutNewRoom GatewayNotifySessionsAboutNewRoom,
) UseCase {
	return useCase{
		gatewayCreateRoom:                 gatewayCreateRoom,
		gatewayNotifySessionsAboutNewRoom: gatewayNotifySessionsAboutNewRoom,
	}
}

type useCase struct {
	gatewayCreateRoom                 GatewayCreateRoom
	gatewayNotifySessionsAboutNewRoom GatewayNotifySessionsAboutNewRoom
}

func (uc useCase) Do(
	ctx context.Context,
	args *create.Args,
) (
	*create.Result,
	error,
) {
	roomEntity, err := uc.
		gatewayCreateRoom.
		Call(
			ctx, args.CreatorUserID, args.RoomName,
		)
	if err != nil {
		return nil, fmt.Errorf("create room: %w", err)
	}

	err = uc.
		gatewayNotifySessionsAboutNewRoom.
		Call(
			ctx, roomEntity,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about new room: %w", err)
	}

	return &create.Result{
		Room: roomEntity,
	}, nil
}
