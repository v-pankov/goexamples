package validator

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *enter.Args) error
}

func NewArgsValidator() ArgsValidator {
	return argsValidator{}
}

type argsValidator struct{}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *enter.Args,
) error {
	if err := args.UserName.Validate(); err != nil {
		return fmt.Errorf("user name: %w", err)
	}
	return nil
}
