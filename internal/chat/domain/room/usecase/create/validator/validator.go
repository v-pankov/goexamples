package validator

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/create"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *create.Args) error
}

var (
	ErrEmptyRoomName       = errors.New("room name is empty")
	ErrNotFoundCreatorUser = errors.New("creator user is not found")
	ErrNotUniqueRoomName   = errors.New("room name is not unique")
)

func NewUseCaseRoomCreateArgsValidator(
	gatewayFindUser GatewayFindUser,
	gatewayFindRoom GatewayFindRoom,
) ArgsValidator {
	return argsValidator{
		gatewayFindUser: gatewayFindUser,
		gatewayFindRoom: gatewayFindRoom,
	}
}

type argsValidator struct {
	gatewayFindUser GatewayFindUser
	gatewayFindRoom GatewayFindRoom
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *create.Args,
) error {
	if len(strings.TrimSpace(args.RoomName.String())) == 0 {
		return ErrEmptyRoomName
	}

	creatorUserEntity, err := v.
		gatewayFindUser.
		Call(
			ctx, args.CreatorUserID,
		)
	if err != nil {
		return fmt.Errorf("find creator user: %w", err)
	}

	if creatorUserEntity == nil {
		return ErrNotFoundCreatorUser
	}

	roomEntity, err := v.
		gatewayFindRoom.
		Call(
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
