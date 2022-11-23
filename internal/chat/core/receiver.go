package core

import "context"

type Receiver interface {
	Receive(context.Context) <-chan []byte
}

func LoopReceiver(
	ctx context.Context,
	receiver Receiver,
	handleFunc func([]byte),
) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case bytes := <-receiver.Receive(ctx):
			handleFunc(bytes)
		}
	}
}
