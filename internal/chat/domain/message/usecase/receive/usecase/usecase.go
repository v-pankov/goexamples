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
	gatewaySessionMessageDeliverer GatewaySessionMessageDeliverer,
) UseCase {
	return useCase{
		gatewaySessionMessageDeliverer: gatewaySessionMessageDeliverer,
	}
}

type useCase struct {
	gatewaySessionMessageDeliverer GatewaySessionMessageDeliverer
}

func (uc useCase) Do(
	ctx context.Context,
	args *receive.Args,
) (
	*receive.Result,
	error,
) {
	err := uc.
		gatewaySessionMessageDeliverer.
		GatewayDeliverMessageToSession(
			ctx,
			args.SessionID,
			args.Message,
		)
	if err != nil {
		return nil, fmt.Errorf("deliver message to session: %w", err)
	}

	return &receive.Result{}, nil
}
