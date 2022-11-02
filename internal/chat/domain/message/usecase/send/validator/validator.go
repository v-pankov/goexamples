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
	ctx context.Context, args *send.Args,
) error {
	authorUserEntity, err := v.
		gatewayFindUser.
		Call(
			ctx, args.AuthorUserID,
		)
	if err != nil {
		return fmt.Errorf("find author user: %w", err)
	}

	if authorUserEntity == nil {
		return ErrNotFoundAuthorUser
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
