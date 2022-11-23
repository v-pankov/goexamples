package app

import "context"

type ErrorHandler interface {
	HandleError(ctx context.Context, err error)
}

type ErrorHandlerFunc func(ctx context.Context, err error)

func (f ErrorHandlerFunc) HandleError(ctx context.Context, err error) {
	f(ctx, err)
}
