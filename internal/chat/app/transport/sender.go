package transport

import "context"

type Sender interface {
	Send(context.Context, []byte) error
}
