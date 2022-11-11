package login

import (
	"context"
	"fmt"
	"strings"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/user/login/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/user/login/model/response"
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

	userEntity, err := uc.gateways.CreateOrFindUser(
		ctx,
		user.Name(requestModel.UserName),
	)
	if err != nil {
		return nil, fmt.Errorf("create or find user: %w", err)
	}

	sessionEntity, err := uc.gateways.CreateSession(
		ctx, userEntity.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return &response.Model{
		SessionID: response.SessionID(sessionEntity.ID),
	}, nil
}
