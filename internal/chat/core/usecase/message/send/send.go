package send

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/core/usecase"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/gateways"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/request"
	"github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/response"
)

type Processor struct {
	Gateways Gateways
}

var _ usecase.Processor[request.Model, response.Model] = Processor{}

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
