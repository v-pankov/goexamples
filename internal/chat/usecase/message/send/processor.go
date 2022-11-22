package send

import (
	"context"
	"fmt"
	"time"
)

type Request struct {
	MessageContents []byte
}

type Response struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       time.Time
}

type Processor interface {
	Process(context.Context, *Request) (*Response, error)
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

func (p processor) Process(ctx context.Context, requestModel *Request) (*Response, error) {
	message, err := p.repository.CreateMessage(
		ctx,
		requestModel.MessageContents,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	return &Response{
		MessageID:       int64(message.ID),
		MessageContents: message.Contents,
		CreatedAt:       message.CreatedAt,
	}, nil
}
