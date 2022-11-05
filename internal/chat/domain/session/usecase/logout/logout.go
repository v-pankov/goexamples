package logout

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/usecase"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/validator"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/values"
)

func ValidateArgsAndRun(
	ctx context.Context,
	useCase usecase.UseCase,
	argsValidator validator.ArgsValidator,
	args *values.Args,
) (
	*values.Result,
	error,
) {
	if err := argsValidator.ValidateArgs(ctx, args); err != nil {
		return nil, fmt.Errorf("validate arguments: %w", err)
	}
	return useCase.Do(ctx, args)
}
