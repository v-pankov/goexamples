package inmem

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/pubsub"
)

type subImpl struct {
	id      int
	channel <-chan []byte
	inmem   *InMem
}

var _ pubsub.Sub = (*subImpl)(nil)

func (s *subImpl) Receive(context.Context) <-chan []byte {
	return s.channel
}

func (s *subImpl) Unsubscribe() error {
	req := &delSubReq{
		id:     s.id,
		result: make(chan *delSubRes),
	}

	s.inmem.delSub <- req
	res := <-req.result

	if res.err != nil {
		return res.err
	}

	return nil
}
