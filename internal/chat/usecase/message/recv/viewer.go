package send

import (
	"context"
	"encoding/json"

	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type ViewModel struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       int64
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
		MessageContents string `json:"contents"`
		CreatedAt       int64  `json:"created_at"`
	}

	bytes, err := json.Marshal(&modelDTO{
		MessageID:       model.MessageID,
		MessageContents: string(model.MessageContents),
		CreatedAt:       model.CreatedAt,
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
