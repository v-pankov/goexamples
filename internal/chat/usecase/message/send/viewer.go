package send

import (
	"context"
	"encoding/json"
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type ViewModel struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       time.Time
}

type Viewer interface {
	View(context.Context, *ViewModel)
}

func NewViewer(
	sender core.Sender,
	errorHandler core.ErrorHandler,
) Viewer {
	return viewer{
		sender:       sender,
		errorHandler: errorHandler,
	}
}

type viewer struct {
	sender       core.Sender
	errorHandler core.ErrorHandler
}

func (v viewer) View(ctx context.Context, model *ViewModel) {
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
		v.errorHandler.HandleError(ctx, err)
		return
	}

	if err := v.sender.Send(ctx, bytes); err != nil {
		v.errorHandler.HandleError(ctx, err)
		return
	}
}
