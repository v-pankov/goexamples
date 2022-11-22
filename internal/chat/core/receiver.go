package core

import "context"

type Receiver interface {
	Receive(context.Context) <-chan []byte
}
