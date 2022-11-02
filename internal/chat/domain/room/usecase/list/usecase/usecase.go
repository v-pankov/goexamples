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
	gatewayGetAllRooms GatewayGetAllRooms,
) UseCase {
	return useCase{
		gatewayGetAllRooms: gatewayGetAllRooms,
	}
}

type useCase struct {
	gatewayGetAllRooms GatewayGetAllRooms
}

func (uc useCase) Do(
	ctx context.Context,
	args *list.Args,
) (
	*list.Result,
	error,
) {
	roomEntities, err := uc.
		gatewayGetAllRooms.
		Call(
			ctx,
		)
	if err != nil {
		return nil, fmt.Errorf("get all rooms: %w", err)
	}

	return &list.Result{
		Rooms: roomEntities,
	}, nil
}
