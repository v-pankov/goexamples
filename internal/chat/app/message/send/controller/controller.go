// Package controller implements user input processing logic.
//
// This package is responsible for
//    * user input format definition
//    * user input parsing
//    * converting user input into send-message-usecase request
//    * interacting with send-message-usecase
package controller

import (
	"context"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/interactor"
	"github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/request"
)

// Controller is a humble object whose role is to parse user input, form
// send message usecase request and interact with business logic.
type Controller struct {
	Interactor interactor.Interactor
}

func (c *Controller) HandleMessage(ctx context.Context, message []byte) error {
	if err := c.Interactor.Interact(
		ctx,
		&request.Model{
			MessageContents: message,
		},
	); err != nil {
		return fmt.Errorf("interact: %w", err)
	}
	return nil
}
