package send

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/transport"

	inmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem"

	sendMsgInmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem/usecase/message/send"
	sendMsgController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/controller"
	sendMsgPresenter "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/presenter"
	sendMsgViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/viewer"
	sendMsgUsecase "github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send"
)

func Run(
	ctx context.Context,
	receiver transport.Receiver,
	sender transport.Sender,
	repo *inmemRepo.InMem,
) {
	controller := sendMsgController.Controller{
		Interactor: sendMsgUsecase.Interactor{
			Processor: sendMsgUsecase.Processor{
				Gateways: sendMsgUsecase.Gateways{
					Repository: sendMsgInmemRepo.Adapter{
						InMem: repo,
					},
				},
			},
			Presenter: sendMsgPresenter.Presenter{
				ModelViewer: sendMsgViewer.Viewer{
					Sender: sender,
				},
			},
		},
	}

	// ignore context cancellation error: it does not matter how it was cancelled
	transport.LoopReceiver(ctx, receiver, func(message []byte) {
		controller.HandleMessage(ctx, message)
	})
}
