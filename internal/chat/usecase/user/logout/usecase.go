package logout

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
	if err := uc.gateways.SessionDeactivator.Deactivate(ctx, request.SessionID); err != nil {
		return nil, fmt.Errorf("deactivate session: %w", err)
	}

	return &Response{}, nil
}
