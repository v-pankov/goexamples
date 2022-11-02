package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseUserEnter interface {
	DoUseCaseUserEnter(
		ctx context.Context,
		args *UseCaseUserEnterArgs,
	) (
		*UseCaseUserEnterResult,
		error,
	)
}

type UseCaseUserEnterArgs struct {
	UserName entity.UserName
}

type UseCaseUserEnterResult struct {
	SessionID entity.SessionID
}

func NewUseCaseUserEnter(
	gatewayUserCreatorFinder UseCaseUserEnterGatewayUserCreatorFinder,
	gatewaySessionCreator UseCaseUserEnterGatewaySessionCreator,
) UseCaseUserEnter {
	return useCaseUserEnter{
		gatewayUserCreatorFinder: gatewayUserCreatorFinder,
		gatewaySessionCreator:    gatewaySessionCreator,
	}
}

type UseCaseUserEnterGatewayUserCreatorFinder interface {
	UseCaseUserEnterGatewayCreateOrFindUser(ctx context.Context, userName entity.UserName) (*entity.User, error)
}

type UseCaseUserEnterGatewaySessionCreator interface {
	UseCaseUserEnterGatewayCreateSession(ctx context.Context, userID entity.UserID) (*entity.Session, error)
}

type useCaseUserEnter struct {
	gatewayUserCreatorFinder UseCaseUserEnterGatewayUserCreatorFinder
	gatewaySessionCreator    UseCaseUserEnterGatewaySessionCreator
}

func (uc useCaseUserEnter) DoUseCaseUserEnter(
	ctx context.Context,
	args *UseCaseUserEnterArgs,
) (
	*UseCaseUserEnterResult,
	error,
) {
	userEntity, err := uc.
		gatewayUserCreatorFinder.
		UseCaseUserEnterGatewayCreateOrFindUser(
			ctx, args.UserName,
		)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", args.UserName, err)
	}

	sessionEntity, err := uc.
		gatewaySessionCreator.
		UseCaseUserEnterGatewayCreateSession(
			ctx, userEntity.ID,
		)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return &UseCaseUserEnterResult{
		SessionID: sessionEntity.ID,
	}, nil
}

type UseCaseUserEnterArgsValidator interface {
	ValidateUseCaseUserEnterArgs(ctx context.Context, args *UseCaseUserEnterArgs) error
}

func NewUseCaseUserEnterArgsValidator() UseCaseUserEnterArgsValidator {
	return useCaseUserEnterArgsValidator{}
}

type useCaseUserEnterArgsValidator struct{}

func (v useCaseUserEnterArgsValidator) ValidateUseCaseUserEnterArgs(
	ctx context.Context, args *UseCaseUserEnterArgs,
) error {
	if err := args.UserName.Validate(); err != nil {
		return fmt.Errorf("user name: %w", err)
	}
	return nil
}
