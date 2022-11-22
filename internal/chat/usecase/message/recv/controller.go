package send

import (
	"context"
	"encoding/json"

	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type Controller struct {
	receiver     core.Receiver
	viewer       Viewer
	errorHandler core.ErrorHandler
}

func NewController(
	receiver core.Receiver,
	viewer Viewer,
	errorHandler core.ErrorHandler,
) *Controller {
	return &Controller{
		receiver:     receiver,
		viewer:       viewer,
		errorHandler: errorHandler,
	}
}

func (c *Controller) Run(ctx context.Context) {
	type input struct {
		MessageID       int64  `json:"id"`
		MessageContents []byte `json:"contents"`
		CreatedAt       int64  `json:"created_at"`
	}

	for message := range c.receiver.Receive(ctx) {
		var input input
		if err := json.Unmarshal(message, &input); err != nil {
			c.errorHandler.HandleError(ctx, err)
			continue
		}

		c.viewer.View(ctx, &ViewModel{
			MessageID:       input.MessageID,
			MessageContents: input.MessageContents,
			CreatedAt:       input.CreatedAt,
		})
	}
}
