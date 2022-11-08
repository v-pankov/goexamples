package login

import (
	"context"
	"fmt"
	"strings"
)

type UseCase interface {
	Do(ctx context.Context, request *Request) (*Response, error)
}

func New(
	gateways Gateways,
) UseCase {
	return useCase{
		gateways: gateways,
	}
}

type useCase struct {
	gateways Gateways
}

func (uc useCase) Do(ctx context.Context, request *Request) (*Response, error) {
	if len(strings.TrimSpace(request.UserName.String())) == 0 {
		return nil, ErrEmptyUserName
	}

	userEntity, err := uc.gateways.UserCreatorFinder.CreateOrFind(ctx, request.UserName)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", request.UserName, err)
	}

	sessionEntity, err := uc.gateways.SessionCreator.Create(ctx, userEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return &Response{
		SessionID: sessionEntity.ID,
	}, nil
}
