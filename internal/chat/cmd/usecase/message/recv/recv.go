package recv

import (
	"context"

	appIO "github.com/vdrpkv/goexamples/internal/chat/app/io"

	recvMsgController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/controller"
	recvMsgViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/viewer"
)

func Run(
	ctx context.Context,
	receiver appIO.Receiver,
	sender appIO.Sender,
) {
	controller := recvMsgController.Controller{
		Viewer: recvMsgViewer.Viewer{
			Sender: sender,
		},
	}

	// ignore context cancellation error: it does not matter how it was cancelled
	_ = appIO.LoopReceiver(ctx, receiver, func(message []byte) {
		controller.HandleMessage(ctx, message)
	})
}
