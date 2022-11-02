package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/delete"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *delete.Args) error
}

var (
	ErrNotFoundRoom = errors.New("room is not found")
)

func New(
	gatewayRoomFinder GatewayRoomFinder,
) ArgsValidator {
	return argsValidator{
		gatewayRoomFinder: gatewayRoomFinder,
	}
}

type argsValidator struct {
	gatewayRoomFinder GatewayRoomFinder
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *delete.Args,
) error {
	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		GatewayFindRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity != nil {
		return ErrNotFoundRoom
	}

	return nil
}
