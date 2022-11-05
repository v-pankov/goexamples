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
) ArgsValidator {
	return argsValidator{
		gatewayFindUser: gatewayFindUser,
	}
}

type argsValidator struct {
	gatewayFindUser GatewayFindUser
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

	return nil
}
