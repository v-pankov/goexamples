package inmem

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/infrastructure/pubsub"
)

type InMem struct {
	nextSubID int
	subs      map[int]chan []byte
	addSub    chan *addSubReq
	delSub    chan *delSubReq
	broadcast chan []byte
}

var _ pubsub.Pub = (*InMem)(nil)

func New() *InMem {
	return &InMem{
		subs:      make(map[int]chan []byte),
		addSub:    make(chan *addSubReq),
		delSub:    make(chan *delSubReq),
		broadcast: make(chan []byte),
	}
}

func (h *InMem) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case addSubReq := <-h.addSub:
			h.processAddSubReq(addSubReq)
		case delSubReq := <-h.delSub:
			h.processDelSubReq(delSubReq)
		case message := <-h.broadcast:
			h.processBroadcastReq(message)
		}
	}
}

func (inmem *InMem) Subscribe() (pubsub.Sub, error) {
	req := &addSubReq{
		result: make(chan *addSubRes),
	}

	inmem.addSub <- req
	res := <-req.result

	if res.err != nil {
		return nil, res.err
	}

	return &subImpl{
		id:      res.id,
		channel: res.channel,
		inmem:   inmem,
	}, nil
}

func (inmem *InMem) Send(_ context.Context, bytes []byte) error {
	inmem.broadcast <- bytes
	return nil
}

type (
	addSubReq struct {
		result chan *addSubRes
	}

	addSubRes struct {
		id      int
		channel <-chan []byte
		err     error
	}

	delSubReq struct {
		id     int
		result chan *delSubRes
	}

	delSubRes struct {
		err error
	}
)

func (inmem *InMem) processAddSubReq(req *addSubReq) {
	var (
		subID   = inmem.getNextSubID()
		channel = make(chan []byte, 256)
	)

	inmem.subs[subID] = channel

	req.result <- &addSubRes{
		id:      subID,
		channel: channel,
	}
}

func (h *InMem) processDelSubReq(req *delSubReq) {
	if channel, ok := h.subs[req.id]; ok {
		delete(h.subs, req.id)
		close(channel)
	}
	req.result <- &delSubRes{}
}

func (h *InMem) processBroadcastReq(message []byte) {
	for subID, channel := range h.subs {
		select {
		case channel <- message:
		default:
			close(channel)
			delete(h.subs, subID)
		}
	}
}

func (h *InMem) getNextSubID() int {
	h.nextSubID++
	return h.nextSubID
}
