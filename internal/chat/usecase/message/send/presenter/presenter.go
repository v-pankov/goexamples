package presenter

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/response"
)

type Presenter interface {
	Present(ctx context.Context, responseModel *response.Model) error
}
