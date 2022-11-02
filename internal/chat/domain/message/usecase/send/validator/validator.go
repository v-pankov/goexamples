package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send"
)

var (
	ErrNotFoundAuthorUser = errors.New("author user is not found")
	ErrNotFoundRoom       = errors.New("room is not found")
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *send.Args) error
}

func New(
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
	ctx context.Context, args *send.Args,
) error {
	if err := args.AuthorUserID.Validate(); err != nil {
		return fmt.Errorf("author user id: %w", err)
	}

	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	authorUserEntity, err := v.
		gatewayUserFinder.
		GatewayFindUser(
			ctx, args.AuthorUserID,
		)
	if err != nil {
		return fmt.Errorf("find author user: %w", err)
	}

	if authorUserEntity == nil {
		return ErrNotFoundAuthorUser
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		GatewayRoomFinder(
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
