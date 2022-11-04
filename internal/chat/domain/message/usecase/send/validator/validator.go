package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send"
)

var (
	ErrNotFoundAuthorUser = errors.New("author user is not found")
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
	authorUserEntity, err := v.repository.FindUser(ctx, args.AuthorUserID)
	if err != nil {
		return fmt.Errorf("find author user: %w", err)
	}

	if authorUserEntity == nil {
		return ErrNotFoundAuthorUser
	}

	return nil
}
