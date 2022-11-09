package send

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/model/response"
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
	messageEntity, err := uc.gateways.MessageCreator.Create(
		ctx,
		session.ID(requestCtx.SessionID),
		requestModel.MessageText,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	if err = uc.gateways.EventCreator.CreateEvent(
		ctx, user.ID(requestCtx.UserID), messageEntity,
	); err != nil {
		return nil, fmt.Errorf("create event: %w", err)
	}

	return &response.Model{}, nil
}
