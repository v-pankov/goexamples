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
	repository Repository,
) UseCase {
	return useCase{
		repository: repository,
	}
}

type useCase struct {
	repository Repository
}

func (uc useCase) Do(
	ctx context.Context,
	args *enter.Args,
) (
	*enter.Result,
	error,
) {
	userEntity, err := uc.repository.CreateOrFindUser(ctx, args.UserName)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", args.UserName, err)
	}

	sessionEntity, err := uc.repository.CreateSession(ctx, userEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return &enter.Result{
		SessionID: sessionEntity.ID,
	}, nil
}
