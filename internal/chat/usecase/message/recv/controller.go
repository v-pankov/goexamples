package send

import (
	"context"
	"encoding/json"

	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type Controller struct {
	viewer       Viewer
	errorHandler core.ErrorHandler
}

func NewController(
	viewer Viewer,
	errorHandler core.ErrorHandler,
) *Controller {
	return &Controller{
		viewer:       viewer,
		errorHandler: errorHandler,
	}
}

func (c *Controller) HandleMessage(ctx context.Context, message []byte) {
	var input struct {
		MessageID       int64  `json:"id"`
		MessageContents []byte `json:"contents"`
		CreatedAt       int64  `json:"created_at"`
	}

	if err := json.Unmarshal(message, &input); err != nil {
		c.errorHandler.HandleError(ctx, err)
		return
	}

	c.viewer.View(ctx, &ViewModel{
		MessageID:       input.MessageID,
		MessageContents: input.MessageContents,
		CreatedAt:       input.CreatedAt,
	})
}
