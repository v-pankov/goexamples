package usecase

import (
	"context"
	"fmt"
)

type UseCase[RequestCtx any, RequestModel any, ResponseModel any] interface {
	Do(
		ctx context.Context,
		requestCtx *RequestCtx,
		requestModel *RequestModel,
	) (
		*ResponseModel,
		error,
	)
}

type Interactor[RequestCtx any, RequestModel any] interface {
	Interact(
		ctx context.Context,
		requestCtx *RequestCtx,
		requestModel *RequestModel,
	) error
}

type Presenter[ResponseModel any] interface {
	Present(
		ctx context.Context,
		responseModel *ResponseModel,
	) error
}

func NewInteractor[RequestCtx any, RequestModel any, ResponseModel any](
	usecase UseCase[RequestCtx, RequestModel, ResponseModel],
	presenter Presenter[ResponseModel],
) Interactor[RequestCtx, RequestModel] {
	return interactor[RequestCtx, RequestModel, ResponseModel]{
		usecase:   usecase,
		presenter: presenter,
	}
}

type interactor[RequestCtx any, RequestModel any, ResponseModel any] struct {
	usecase   UseCase[RequestCtx, RequestModel, ResponseModel]
	presenter Presenter[ResponseModel]
}

func (i interactor[RequestCtx, RequestModel, ResponseModel]) Interact(
	ctx context.Context,
	requestCtx *RequestCtx,
	requestModel *RequestModel,
) error {
	responseModel, err := i.usecase.Do(ctx, requestCtx, requestModel)
	if err != nil {
		return fmt.Errorf("usecase: %w", err)
	}

	if err := i.presenter.Present(ctx, responseModel); err != nil {
		return fmt.Errorf("present: %w", err)
	}

	return nil
}
