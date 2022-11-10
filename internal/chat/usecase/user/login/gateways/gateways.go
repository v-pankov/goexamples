package gateways

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type UserCreator interface {
	CreateUser(ctx context.Context, userName user.Name) (user.ID, error)
}

type UserFinder interface {
	FindUser(ctx context.Context, userName user.Name) (*user.Entity, error)
}

func CreateOrFindUser(
	ctx context.Context,
	userCreator UserCreator,
	userFinder UserFinder,
	userName user.Name,
) (
	user.ID,
	error,
) {
	userEntity, err := userFinder.FindUser(ctx, userName)
	if err != nil {
		return "", fmt.Errorf("find user: %w", err)
	}

	if userEntity != nil {
		return userEntity.ID, nil
	}

	userID, err := userCreator.CreateUser(ctx, userName)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	return userID, nil
}
