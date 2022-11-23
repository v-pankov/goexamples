package interactor

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/request"
)

type Interactor interface {
	Interact(ctx context.Context, requestModel *request.Model) error
}
