package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *enter.Args,
	) (
		*enter.Result,
		error,
	)
}

func New(
	gatewayCreateOrFindUser GatewayCreateOrFindUser,
	gatewayCreateSession GatewayCreateSession,
) UseCase {
	return useCase{
		gatewayCreateOrFindUser: gatewayCreateOrFindUser,
		gatewayCreateSession:    gatewayCreateSession,
	}
}

type useCase struct {
	gatewayCreateOrFindUser GatewayCreateOrFindUser
	gatewayCreateSession    GatewayCreateSession
}

func (uc useCase) Do(
	ctx context.Context,
	args *enter.Args,
) (
	*enter.Result,
	error,
) {
	userEntity, err := uc.
		gatewayCreateOrFindUser.
		Call(
			ctx, args.UserName,
		)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", args.UserName, err)
	}

	sessionEntity, err := uc.
		gatewayCreateSession.
		Call(
			ctx, userEntity.ID,
		)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return &enter.Result{
		SessionID: sessionEntity.ID,
	}, nil
}
