package authenticate

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/session/authenticate/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/session/authenticate/model/response"
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
	sessionEntity, err := uc.gateways.FindSessionEntity(
		ctx, session.ID(requestModel.SessionID),
	)
	if err != nil {
		return nil, fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return nil, ErrSessionNotFound
	}

	if !sessionEntity.Active {
		return nil, ErrSessionNotActive
	}

	return &response.Model{
		UserID: response.UserID(sessionEntity.UserID),
	}, nil
}
