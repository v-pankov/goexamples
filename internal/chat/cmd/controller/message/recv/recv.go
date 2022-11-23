package recv

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app/controller"
	"github.com/vdrpkv/goexamples/internal/chat/app/transport"

	recvMsgController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/controller"
	recvMsgViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/viewer"
)

func New(
	ctx context.Context,
	sender transport.Sender,
) controller.Controller {
	return recvMsgController.Controller{
		Viewer: recvMsgViewer.Viewer{
			Sender: sender,
		},
	}
}
