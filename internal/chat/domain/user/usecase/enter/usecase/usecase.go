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
	gatewayUserCreatorFinder GatewayUserCreatorFinder,
	gatewaySessionCreator GatewaySessionCreator,
) UseCase {
	return useCase{
		gatewayUserCreatorFinder: gatewayUserCreatorFinder,
		gatewaySessionCreator:    gatewaySessionCreator,
	}
}

type useCase struct {
	gatewayUserCreatorFinder GatewayUserCreatorFinder
	gatewaySessionCreator    GatewaySessionCreator
}

func (uc useCase) Do(
	ctx context.Context,
	args *enter.Args,
) (
	*enter.Result,
	error,
) {
	userEntity, err := uc.
		gatewayUserCreatorFinder.
		GatewayCreateOrFindUser(
			ctx, args.UserName,
		)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", args.UserName, err)
	}

	sessionEntity, err := uc.
		gatewaySessionCreator.
		GatewayCreateSession(
			ctx, userEntity.ID,
		)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return &enter.Result{
		SessionID: sessionEntity.ID,
	}, nil
}
