package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/list"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *list.Args,
	) (
		*list.Result,
		error,
	)
}

func New(
	gatewayRoomGetter GatewayRoomGetter,
) UseCase {
	return useCase{
		gatewayRoomGetter: gatewayRoomGetter,
	}
}

type useCase struct {
	gatewayRoomGetter GatewayRoomGetter
}

func (uc useCase) Do(
	ctx context.Context,
	args *list.Args,
) (
	*list.Result,
	error,
) {
	roomEntities, err := uc.
		gatewayRoomGetter.
		GatewayGetAllRooms(
			ctx,
		)
	if err != nil {
		return nil, fmt.Errorf("get all rooms: %w", err)
	}

	return &list.Result{
		Rooms: roomEntities,
	}, nil
}
