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
	gatewayFindRoom GatewayFindRoom,
) ArgsValidator {
	return argsValidator{
		gatewayFindRoom: gatewayFindRoom,
	}
}

type argsValidator struct {
	gatewayFindRoom GatewayFindRoom
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *delete.Args,
) error {
	roomEntity, err := v.
		gatewayFindRoom.
		Call(
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
