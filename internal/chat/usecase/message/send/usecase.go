package send

import (
	"context"
	"fmt"
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
	messageEntity, err := uc.gateways.MessageCreator.Create(
		ctx,
		request.SessionID,
		request.MessageText,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	if err = uc.gateways.MessageBroadcaster.Broadcast(ctx, messageEntity); err != nil {
		return nil, fmt.Errorf("broadcast message: %w", err)
	}

	return &Response{}, nil
}
