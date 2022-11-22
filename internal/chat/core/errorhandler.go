package core

import (
	"context"
)

type ErrorHandler interface {
	HandleError(context.Context, error)
}

type ErrorHandlerFunc func(context.Context, error)

func (f ErrorHandlerFunc) HandleError(ctx context.Context, err error) {
	f(ctx, err)
}
