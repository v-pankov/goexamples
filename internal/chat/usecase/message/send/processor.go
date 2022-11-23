package send

import (
	"context"
	"fmt"
	"time"
)

type RequestModel struct {
	MessageContents []byte
}

type ResponseModel struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       time.Time
}

type Processor interface {
	Process(context.Context, *RequestModel) (*ResponseModel, error)
}

type processor struct {
	repository Repository
}

func NewProcessor(
	repository Repository,
) Processor {
	return processor{
		repository: repository,
	}
}

func (p processor) Process(ctx context.Context, requestModel *RequestModel) (*ResponseModel, error) {
	message, err := p.repository.CreateMessage(
		ctx,
		requestModel.MessageContents,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	return &ResponseModel{
		MessageID:       int64(message.ID),
		MessageContents: message.Contents,
		CreatedAt:       message.CreatedAt,
	}, nil
}
