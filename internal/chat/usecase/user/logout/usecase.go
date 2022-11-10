package logout

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/model/response"
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
	if err := uc.gateways.DeleteSession(ctx, session.ID(requestCtx.SessionID)); err != nil {
		return nil, fmt.Errorf("delete session: %w", err)
	}

	return &response.Model{}, nil
}
