package websocket

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/vdrpkv/goexamples/internal/chat/app"
	"github.com/vdrpkv/goexamples/internal/chat/app/controller"
	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/pubsub"
	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/websocket"

	inmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem"

	msgRecvController "github.com/vdrpkv/goexamples/internal/chat/cmd/controller/message/recv"
	msgSendController "github.com/vdrpkv/goexamples/internal/chat/cmd/controller/message/send"
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

	for _, loop := range []controller.Loop{
		{
			Receiver: wsHandler,
			Controller: msgSendController.New(
				ctx, pub, inmemRepo,
			),
		},
		{
			Receiver: sub,
			Controller: msgRecvController.New(
				ctx, wsHandler,
			),
		},
	} {
		wg.Add(1)
		go func(loop controller.Loop) {
			// ignore context cancellation error
			_ = loop.Run(ctx, app.ErrorHandlerFunc(
				func(ctx context.Context, err error) {
					log.Println(err)
				}),
			)
			wg.Done()
		}(loop)
	}
}
