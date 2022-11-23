package send

import (
	"context"
)

type Controller struct {
	interactor Interactor
}

func NewController(
	interactor Interactor,
) *Controller {
	return &Controller{
		interactor: interactor,
	}
}

func (c *Controller) HandleMessage(ctx context.Context, message []byte) {
	c.interactor.Interact(
		ctx,
		&Request{
			MessageContents: message,
		},
	)
}
