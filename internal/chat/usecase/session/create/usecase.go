package create

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/session/create/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/session/create/model/response"
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
	sessionEntity, err := uc.gateways.CreateNewSessionEntity(
		ctx, user.ID(requestModel.UserID),
	)
	if err != nil {
		return nil, fmt.Errorf("create new user entity: %w", err)
	}

	if err := uc.gateways.CreateNewSessionEvent(
		ctx, sessionEntity,
	); err != nil {
		return nil, fmt.Errorf("create new user event: %w", err)
	}

	return &response.Model{
		SessionID: response.SessionID(sessionEntity.ID),
	}, nil
}
