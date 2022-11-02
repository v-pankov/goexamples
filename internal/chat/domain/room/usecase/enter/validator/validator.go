package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/enter"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *enter.Args) error
}

var (
	ErrNotFoundSession = errors.New("session is not found")
	ErrNotFoundRoom    = errors.New("room is not found")
)

func New(
	gatewaySessionFinder GatewaySessionFinder,
	gatewayRoomFinder GatewayRoomFinder,
) ArgsValidator {
	return argsValidator{
		gatewaySessionFinder: gatewaySessionFinder,
		gatewayRoomFinder:    gatewayRoomFinder,
	}
}

type argsValidator struct {
	gatewaySessionFinder GatewaySessionFinder
	gatewayRoomFinder    GatewayRoomFinder
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *enter.Args,
) error {
	if err := args.SessionID.Validate(); err != nil {
		return fmt.Errorf("session id: %w", err)
	}

	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	sessionEntity, err := v.
		gatewaySessionFinder.
		GatewayFindSession(
			ctx, args.SessionID,
		)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrNotFoundSession
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		GatewayFindRoom(
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
