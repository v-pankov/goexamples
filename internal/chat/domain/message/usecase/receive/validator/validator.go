package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/receive"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *receive.Args) error
}

var (
	ErrNotFoundSession = errors.New("session is not found")
)

func New(
	gatewaySessionFinder GatewaySessionFinder,
) ArgsValidator {
	return argsValidator{
		gatewaySessionFinder: gatewaySessionFinder,
	}
}

type argsValidator struct {
	gatewaySessionFinder GatewaySessionFinder
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *receive.Args,
) error {
	if err := args.SessionID.Validate(); err != nil {
		return fmt.Errorf("session id: %w", err)
	}

	if err := args.Message.Validate(); err != nil {
		return fmt.Errorf("message: %w", err)
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

	return nil
}
