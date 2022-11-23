package recv

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/transport"

	recvMsgController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/controller"
	recvMsgViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/viewer"
)

func Run(
	ctx context.Context,
	receiver transport.Receiver,
	sender transport.Sender,
) {
	controller := recvMsgController.Controller{
		Viewer: recvMsgViewer.Viewer{
			Sender: sender,
		},
	}

	// ignore context cancellation error: it does not matter how it was cancelled
	_ = transport.LoopReceiver(ctx, receiver, func(message []byte) {
		controller.HandleMessage(ctx, message)
	})
}
