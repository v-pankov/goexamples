package io

import "context"

type Sender interface {
	Send(context.Context, []byte) error
}
