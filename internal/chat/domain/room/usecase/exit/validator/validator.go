package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/exit"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *exit.Args) error
}

var (
	ErrNotFoundSession = errors.New("session is not found")
	ErrNotFoundRoom    = errors.New("room is not found")
)

func New(
	gatewayFindSession GatewayFindSession,
	gatewayFindRoom GatewayFindRoom,
) ArgsValidator {
	return argsValidator{
		gatewayFindSession: gatewayFindSession,
		gatewayFindRoom:    gatewayFindRoom,
	}
}

type argsValidator struct {
	gatewayFindSession GatewayFindSession
	gatewayFindRoom    GatewayFindRoom
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *exit.Args,
) error {
	if err := args.SessionID.Validate(); err != nil {
		return fmt.Errorf("session id: %w", err)
	}

	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	sessionEntity, err := v.
		gatewayFindSession.
		Call(
			ctx, args.SessionID,
		)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrNotFoundSession
	}

	roomEntity, err := v.
		gatewayFindRoom.
		Call(
			ctx, args.RoomID,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity == nil {
		return ErrNotFoundRoom
	}

	return nil
}
