package send

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/transport"

	coreUsecase "github.com/vdrpkv/goexamples/internal/chat/core/usecase"

	appRepoInmem "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem"
	appRepoInmemUsecaseAdapter "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem/usecase/message/send"

	appController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/controller"
	appPresenter "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/presenter"
	appViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/viewer"

	usecase "github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send"
	usecaseRequest "github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/request"
	usecaseResponse "github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send/response"
)

func Run(
	ctx context.Context,
	receiver transport.Receiver,
	sender transport.Sender,
	repo *appRepoInmem.InMem,
) {
	controller := appController.Controller{
		Interactor: coreUsecase.NewInteractor[
			usecaseRequest.Model,
			usecaseResponse.Model,
		](
			usecase.Processor{
				Gateways: usecase.Gateways{
					Repository: appRepoInmemUsecaseAdapter.Adapter{
						InMem: repo,
					},
				},
			},
			appPresenter.Presenter{
				ModelViewer: appViewer.Viewer{
					Sender: sender,
				},
			},
		),
	}

	// ignore context cancellation error: it does not matter how it was cancelled
	transport.LoopReceiver(ctx, receiver, func(message []byte) {
		controller.HandleMessage(ctx, message)
	})
}
