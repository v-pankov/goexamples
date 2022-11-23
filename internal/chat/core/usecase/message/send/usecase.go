package usecase

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/gateways"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/presenter"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/processor"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/request"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/response"
)

type Interactor struct {
	Processor processor.Processor
	Presenter presenter.Presenter
}

func (i Interactor) Interact(
	ctx context.Context,
	requestModel *request.Model,
) error {
	rsp, err := i.Processor.Process(ctx, requestModel)
	if err != nil {
		return fmt.Errorf("process request: %w", err)
	}

	if err := i.Presenter.Present(ctx, rsp); err != nil {
		return fmt.Errorf("present response: %w", err)
	}

	return nil
}

type Processor struct {
	Gateways Gateways
}

func (p Processor) Process(
	ctx context.Context,
	requestModel *request.Model,
) (*response.Model, error) {
	message, err := p.Gateways.Repository.CreateMessage(
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

type Gateways struct {
	Repository gateways.Repository
}
