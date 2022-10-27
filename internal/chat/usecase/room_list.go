package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseRoomList interface {
	DoUseCaseRoomList(
		ctx context.Context,
		args *UseCaseRoomListArgs,
	) (
		*UseCaseRoomListResult,
		error,
	)
}

type UseCaseRoomListArgs struct {
}

type UseCaseRoomListResult struct {
	Rooms []entity.Room
}

func NewUseCaseRoomList(
	gatewayRoomGetter UseCaseRoomListGatewayRoomGetter,
) useCaseRoomList {
	return useCaseRoomList{
		gatewayRoomGetter: gatewayRoomGetter,
	}
}

type UseCaseRoomListGatewayRoomGetter interface {
	UseCaseRoomListGatewayGetAllRooms(
		ctx context.Context,
	) (
		[]entity.Room,
		error,
	)
}

type useCaseRoomList struct {
	gatewayRoomGetter UseCaseRoomListGatewayRoomGetter
}

func (uc useCaseRoomList) DoUseCaseListRoom(
	ctx context.Context,
	args *UseCaseRoomListArgs,
) (
	*UseCaseRoomListResult,
	error,
) {
	roomEntities, err := uc.
		gatewayRoomGetter.
		UseCaseRoomListGatewayGetAllRooms(
			ctx,
		)
	if err != nil {
		return nil, fmt.Errorf("get all rooms: %w", err)
	}

	return &UseCaseRoomListResult{
		Rooms: roomEntities,
	}, nil
}

type UseCaseRoomListArgsValidator interface {
	ValidateUseCaseRoomListArgs(ctx context.Context, args *UseCaseRoomListArgs) error
}

func NewUseCaseRoomListArgsValidator() UseCaseRoomListArgsValidator {
	return useCaseRoomListArgsValidator{}
}

type useCaseRoomListArgsValidator struct {
}

func (v useCaseRoomListArgsValidator) ValidateUseCaseRoomListArgs(
	ctx context.Context, args *UseCaseRoomListArgs,
) error {
	return nil
}
