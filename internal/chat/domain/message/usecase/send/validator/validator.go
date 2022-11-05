package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send"
)

var (
	ErrNotFoundAuthorUserSession = errors.New("author user session is not found")
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *send.Args) error
}

func New(
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
	ctx context.Context, args *send.Args,
) error {
	authorUserSessionEntity, err := v.repository.FindSession(ctx, args.AuthorUserSessionID)
	if err != nil {
		return fmt.Errorf("find author user session: %w", err)
	}

	if authorUserSessionEntity == nil {
		return ErrNotFoundAuthorUserSession
	}

	return nil
}
