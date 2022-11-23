package send

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app/controller"
	"github.com/vdrpkv/goexamples/internal/chat/app/transport"

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

func New(
	ctx context.Context,
	sender transport.Sender,
	repo *appRepoInmem.InMem,
) controller.Controller {
	return appController.Controller{
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
}
