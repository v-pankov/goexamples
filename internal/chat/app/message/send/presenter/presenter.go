package presenter

import (
	"context"
	"fmt"
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/response"
)

type Presenter struct {
	ModelViewer ModelViewer
}

func (p Presenter) Present(ctx context.Context, rsp *response.Model) error {
	if err := p.ModelViewer.ViewModel(
		ctx,
		&ViewModel{
			MessageID:       rsp.MessageID,
			MessageContents: rsp.MessageContents,
			CreatedAt:       rsp.CreatedAt,
		},
	); err != nil {
		return fmt.Errorf("view model: %w:", err)
	}
	return nil
}

type ViewModel struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       time.Time
}

type ModelViewer interface {
	ViewModel(context.Context, *ViewModel) error
}
