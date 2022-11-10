package login

import (
	"context"
	"fmt"
	"strings"

	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/user/create/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/user/create/model/response"
)

type UseCase interface {
	Do(
		ctx context.Context,
		requestCtx *request.Context,
		requestModel *request.Model,
	) (
		*response.Model,
		error,
	)
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

func (uc useCase) Do(
	ctx context.Context,
	requestCtx *request.Context,
	requestModel *request.Model,
) (
	*response.Model,
	error,
) {
	if len(strings.TrimSpace(requestModel.UserName.String())) == 0 {
		return nil, ErrEmptyUserName
	}

	userEntity, err := uc.gateways.CreateNewUserEntity(
		ctx, user.Name(requestModel.UserName),
	)
	if err != nil {
		return nil, fmt.Errorf("create new user entity: %w", err)
	}

	if err := uc.gateways.CreateNewUserEvent(
		ctx, userEntity,
	); err != nil {
		return nil, fmt.Errorf("create new user event: %w", err)
	}

	return &response.Model{
		UserID: response.UserID(userEntity.ID),
	}, nil
}
