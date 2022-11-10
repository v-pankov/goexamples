package create

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/model/response"
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
	messageEntity, err := uc.gateways.CreateMessage(
		ctx,
		session.ID(requestCtx.SessionID),
		requestModel.MessageText,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	return &response.Model{
		Message: response.Message{
			UserID:      response.UserID(requestCtx.UserID),
			SessionID:   response.SessionID(messageEntity.SessionID),
			MessageID:   response.MessageID(messageEntity.ID),
			MessageText: response.MessageText(messageEntity.Text),
			CreatedAt:   messageEntity.CreatedAt,
		},
	}, nil
}
