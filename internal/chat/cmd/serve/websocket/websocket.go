package websocket

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/pubsub"
	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/websocket"

	inmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem"

	usecaseMsgRecv "github.com/vdrpkv/goexamples/internal/chat/cmd/usecase/message/recv"
	usecaseMsgSend "github.com/vdrpkv/goexamples/internal/chat/cmd/usecase/message/send"
)

func Handler(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	wg *sync.WaitGroup,
	pub pubsub.Pub,
	inmemRepo *inmemRepo.InMem,
) {
	wsHandler := websocket.New()

	if err := wsHandler.Serve(ctx, w, r, websocket.Text); err != nil {
		log.Println(err)
		return
	}

	sub, err := pub.Subscribe()
	if err != nil {
		log.Println(err)
		return
	}

	wg.Add(1)
	go func() {
		usecaseMsgSend.Run(
			ctx, wsHandler, pub, inmemRepo,
		)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		usecaseMsgRecv.Run(
			ctx, sub, wsHandler,
		)
		wg.Done()
	}()
}
