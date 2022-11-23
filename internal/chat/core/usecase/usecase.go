package usecase

import (
	"context"
	"fmt"
)

type Interactor[RequestModel any] interface {
	Interact(ctx context.Context, requestModel *RequestModel) error
}

type Presenter[ResponseModel any] interface {
	Present(ctx context.Context, responseModel *ResponseModel) error
}

type Processor[RequestModel any, ResponseModel any] interface {
	Process(ctx context.Context, requestModel *RequestModel) (*ResponseModel, error)
}

// NewInteractor erases response model information for caller
func NewInteractor[RequestModel any, ResponseModel any](
	processor Processor[RequestModel, ResponseModel],
	presenter Presenter[ResponseModel],
) Interactor[RequestModel] {
	return interactor[RequestModel, ResponseModel]{
		processor: processor,
		presenter: presenter,
	}
}

type interactor[RequestModel any, ResponseModel any] struct {
	processor Processor[RequestModel, ResponseModel]
	presenter Presenter[ResponseModel]
}

func (i interactor[RequestModel, ResponseModel]) Interact(ctx context.Context, requestModel *RequestModel) error {
	rsp, err := i.processor.Process(ctx, requestModel)
	if err != nil {
		return fmt.Errorf("process request: %w", err)
	}

	if err := i.presenter.Present(ctx, rsp); err != nil {
		return fmt.Errorf("present response: %w", err)
	}

	return nil
}
