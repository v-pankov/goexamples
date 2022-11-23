package processor

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/request"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/response"
)

// Processor processes send message requests.
type Processor interface {
	Process(ctx context.Context, requestModel *request.Model) (*response.Model, error)
}
