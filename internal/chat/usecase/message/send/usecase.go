package send

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/model/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/model/response"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"
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
	if err := uc.gateways.SendMessage(
		ctx,
		user.ID(requestCtx.UserID),
		&message.Entity{
			Entity: entity.Entity{
				CreatedAt: requestModel.CreatedAt,
			},
			ID:        message.ID(requestModel.MessageID),
			SessionID: session.ID(requestModel.SessionID),
			Text:      string(requestModel.MessageText),
		},
	); err != nil {
		return nil, fmt.Errorf("send message: %w", err)
	}

	return nil, nil
}
