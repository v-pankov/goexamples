package authenticate

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
	sessionEntity, err := uc.gateways.SessionFinder.Find(ctx, request.SessionID)
	if err != nil {
		return nil, fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return nil, ErrSessionNotFound
	}

	if !sessionEntity.Active {
		return nil, ErrSessionNotActive
	}

	return &Response{}, nil
}
