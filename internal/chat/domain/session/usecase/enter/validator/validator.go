package validator

import (
	"context"
	"errors"
	"strings"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/enter"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *enter.Args) error
}

var (
	ErrEmptyUserName = errors.New("user name is empty")
)

func NewArgsValidator() ArgsValidator {
	return argsValidator{}
}

type argsValidator struct{}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *enter.Args,
) error {
	if len(strings.TrimSpace(args.UserName.String())) == 0 {
		return ErrEmptyUserName
	}

	return nil
}
