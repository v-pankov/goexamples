package viewer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/controller"

	appIO "github.com/vdrpkv/goexamples/internal/chat/app/io"
)

type Viewer struct {
	Sender appIO.Sender
}

var _ controller.ModelViewer = Viewer{}

func (v Viewer) ViewModel(ctx context.Context, model *controller.ViewModel) error {
	type modelDTO struct {
		MessageID       int64  `json:"id"`
		MessageContents string `json:"contents"`
		CreatedAt       int64  `json:"created_at"`
	}

	bytes, err := json.Marshal(&modelDTO{
		MessageID:       model.MessageID,
		MessageContents: string(model.MessageContents),
		CreatedAt:       model.CreatedAt,
	})

	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	if err := v.Sender.Send(ctx, bytes); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
