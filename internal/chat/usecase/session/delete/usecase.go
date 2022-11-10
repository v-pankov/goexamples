package delete

import (
	"context"
	"fmt"
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/session/delete/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/session/delete/model/response"
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
	sesssionEntity, err := uc.gateways.FindSessionEntity(
		ctx, session.ID(requestModel.SessionID),
	)
	if err != nil {
		return nil, fmt.Errorf("find session entity: %w", err)
	}

	if sesssionEntity == nil {
		return nil, ErrSessionNotFound
	}

	if !sesssionEntity.Active {
		return nil, ErrSessionDeleted
	}

	sesssionEntity.Active = false
	sesssionEntity.UpdatedAt = time.Now()

	if err := uc.gateways.UpdateSessionEntity(
		ctx, sesssionEntity,
	); err != nil {
		return nil, fmt.Errorf("update session entity: %w", err)
	}

	if err := uc.gateways.CreateSessionDeletedEvent(
		ctx, sesssionEntity,
	); err != nil {
		return nil, fmt.Errorf("create session deleted event: %w", err)
	}

	return &response.Model{}, nil
}
