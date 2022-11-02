package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/receive"
)

type UseCase interface {
	Do(
		ctx context.Context,
		args *receive.Args,
	) (
		*receive.Result,
		error,
	)
}

func New(
	gatewayDeliverMessageToSession GatewayDeliverMessageToSession,
) UseCase {
	return useCase{
		gatewayDeliverMessageToSession: gatewayDeliverMessageToSession,
	}
}

type useCase struct {
	gatewayDeliverMessageToSession GatewayDeliverMessageToSession
}

func (uc useCase) Do(
	ctx context.Context,
	args *receive.Args,
) (
	*receive.Result,
	error,
) {
	err := uc.
		gatewayDeliverMessageToSession.
		Call(
			ctx,
			args.SessionID,
			args.Message,
		)
	if err != nil {
		return nil, fmt.Errorf("deliver message to session: %w", err)
	}

	return &receive.Result{}, nil
}
