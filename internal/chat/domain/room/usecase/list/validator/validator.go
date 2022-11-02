package validator

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room/usecase/list"
)

type ArgsValidator interface {
	ValidateArgs(ctx context.Context, args *list.Args) error
}

func NewUseCaseRoomListArgsValidator() ArgsValidator {
	return argsValidator{}
}

type argsValidator struct {
}

func (v argsValidator) ValidateArgs(
	ctx context.Context, args *list.Args,
) error {
	return nil
}
