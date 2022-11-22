package send

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type Controller struct {
	receiver   core.Receiver
	interactor Interactor
}

func NewController(
	receiver core.Receiver,
	interactor Interactor,
) *Controller {
	return &Controller{
		receiver:   receiver,
		interactor: interactor,
	}
}

func (c *Controller) Run(ctx context.Context) {
	for message := range c.receiver.Receive(ctx) {
		c.interactor.Interact(
			ctx,
			&Request{
				MessageContents: message,
			},
		)
	}
}
