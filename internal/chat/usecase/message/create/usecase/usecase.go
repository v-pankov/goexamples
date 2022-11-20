package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
	"github.com/vdrpkv/goexamples/internal/chat/usecase"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/usecase/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/create/usecase/response"
)

func (uc useCase) Do(
	ctx context.Context,
	requestCtx *request.Context,
	requestModel *request.Model,
) (
	*response.Model,
	error,
) {
	message, err := uc.repository.CreateMessage(
		ctx,
		requestModel.MessageContents,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	return &response.Model{
		MessageID:       int64(message.ID),
		MessageContents: message.Contents,
		CreatedAt:       message.CreatedAt,
	}, nil
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

type useCase struct {
	repository Repository
}

func New(
	repository Repository,
) usecase.UseCase[request.Context, request.Model, response.Model] {
	return useCase{
		repository: repository,
	}
}
