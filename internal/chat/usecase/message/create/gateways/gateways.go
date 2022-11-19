package gateways

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/usecase"
)

func New(
	repository Repository,
) usecase.Gateways {
	return gateways{
		repository: repository,
	}
}

func (g gateways) CreateMessage(
	ctx context.Context,
	messageContents entity.MessageContents,
) (
	*entity.Message,
	error,
) {
	message, err := g.repository.CreateMessage(ctx, messageContents)
	if err != nil {
		return nil, fmt.Errorf("repository: %w", err)
	}
	return message, nil
}

type gateways struct {
	repository Repository
}

type Repository interface {
	CreateMessage(
		ctx context.Context,
		messageContents entity.MessageContents,
	) (
		*entity.Message,
		error,
	)
}
