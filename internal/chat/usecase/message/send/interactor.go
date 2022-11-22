package send

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/core"
)

type Interactor interface {
	Interact(context.Context, *Request)
}

func NewInteractor(
	processor Processor,
	presenter Presenter,
	errorHandler core.ErrorHandler,
) Interactor {
	return interactor{
		processor:    processor,
		presenter:    presenter,
		errorHandler: errorHandler,
	}
}

type interactor struct {
	processor    Processor
	presenter    Presenter
	errorHandler core.ErrorHandler
}

func (i interactor) Interact(ctx context.Context, req *Request) {
	rsp, err := i.processor.Process(ctx, req)
	if err != nil {
		i.errorHandler.HandleError(ctx, err)
		return
	}
	i.presenter.Present(ctx, rsp)
}
