package viewer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/app/message/send/presenter"
	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type Viewer struct {
	Sender core.Sender
}

func (v Viewer) ViewModel(ctx context.Context, model *presenter.ViewModel) error {
	type modelDTO struct {
		MessageID       int64  `json:"id"`
		MessageContents []byte `json:"contents"`
		CreatedAt       int64  `json:"created_at"`
	}

	bytes, err := json.Marshal(&modelDTO{
		MessageID:       model.MessageID,
		MessageContents: model.MessageContents,
		CreatedAt:       model.CreatedAt.Unix(),
	})
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	if err := v.Sender.Send(ctx, bytes); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
