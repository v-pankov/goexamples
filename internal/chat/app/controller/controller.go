package controller

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app"
	"github.com/vdrpkv/goexamples/internal/chat/app/transport"
)

type Controller interface {
	HandleMessage(ctx context.Context, message []byte) error
}

type Loop struct {
	Receiver   transport.Receiver
	Controller Controller
}

func (loop Loop) Run(ctx context.Context, errorHandler app.ErrorHandler) error {
	return transport.LoopReceiver(ctx, loop.Receiver, func(message []byte) {
		if err := loop.Controller.HandleMessage(ctx, message); err != nil {
			errorHandler.HandleError(ctx, err)
		}
	})
}
