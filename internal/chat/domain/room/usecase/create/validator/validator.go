package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/create"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *create.Args) error
}

var (
	ErrNotFoundCreatorUser = errors.New("creator user is not found")
	ErrNotUniqueRoomName   = errors.New("room name is not unique")
)

func NewUseCaseRoomCreateArgsValidator(
	gatewayUserFinder GatewayUserFinder,
	gatewayRoomFinder GatewayRoomFinder,
) ArgsValidator {
	return argsValidator{
		gatewayUserFinder: gatewayUserFinder,
		gatewayRoomFinder: gatewayRoomFinder,
	}
}

type argsValidator struct {
	gatewayUserFinder GatewayUserFinder
	gatewayRoomFinder GatewayRoomFinder
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *create.Args,
) error {
	if err := args.CreatorUserID.Validate(); err != nil {
		return fmt.Errorf("creator user id: %w", err)
	}

	if err := args.RoomName.Validate(); err != nil {
		return fmt.Errorf("room name: %w", err)
	}

	creatorUserEntity, err := v.
		gatewayUserFinder.
		GatewayFindUser(
			ctx, args.CreatorUserID,
		)
	if err != nil {
		return fmt.Errorf("find creator user: %w", err)
	}

	if creatorUserEntity == nil {
		return ErrNotFoundCreatorUser
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		GatewayFindRoom(
			ctx, args.RoomName,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity != nil {
		return ErrNotUniqueRoomName
	}

	return nil
}
