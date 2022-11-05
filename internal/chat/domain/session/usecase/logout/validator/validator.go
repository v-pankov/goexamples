package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *logout.Args) error
}

var (
	ErrNotFoundSession = errors.New("session  is not found")
)

func NewArgsValidator(
	repository Repository,
) ArgsValidator {
	return argsValidator{
		repository: repository,
	}
}

type argsValidator struct {
	repository Repository
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *logout.Args,
) error {
	sessionEntity, err := v.repository.FindSession(ctx, args.SessionID)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrNotFoundSession
	}

	return nil
}