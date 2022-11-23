package cmd

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"

	appIO "github.com/vdrpkv/goexamples/internal/chat/app/io"

	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/pubsub"
	"github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/websocket"

	inmemPubSub "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/pubsub/inmem"
	inmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem"

	sendMsgInmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem/message/send"
	sendMsgController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/controller"
	sendMsgPresenter "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/presenter"
	sendMsgViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/send/viewer"
	sendMsgUsecase "github.com/vdrpkv/goexamples/internal/chat/core/usecase/message/send"

	recvMsgController "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/controller"
	recvMsgViewer "github.com/vdrpkv/goexamples/internal/chat/app/usecase/message/recv/viewer"
)

var addr = flag.String("addr", ":8080", "http service address")

func Run() {
	flag.Parse()

	ctx := context.Background()

	pub := inmemPubSub.New()
	go pub.Run(ctx)

	inmemRepo := inmemRepo.New()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(ctx, pub, inmemRepo, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./assets/chat/home.html")
}

// serveWs handles websocket requests from the peer.
func serveWs(
	ctx context.Context,
	pub pubsub.Pub,
	inmemRepo *inmemRepo.InMem,
	w http.ResponseWriter,
	r *http.Request,
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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		setupSendMessageUsecase(
			ctx, wsHandler, pub, inmemRepo,
		)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		setupRecvMessageUsecase(
			ctx, sub, wsHandler,
		)
		wg.Done()
	}()
}

func setupSendMessageUsecase(
	ctx context.Context,
	receiver appIO.Receiver,
	sender appIO.Sender,
	repo *inmemRepo.InMem,
) {
	controller := sendMsgController.Controller{
		Interactor: sendMsgUsecase.Interactor{
			Processor: sendMsgUsecase.Processor{
				Gateways: sendMsgUsecase.Gateways{
					Repository: sendMsgInmemRepo.Adapter{
						InMem: repo,
					},
				},
			},
			Presenter: sendMsgPresenter.Presenter{
				ModelViewer: sendMsgViewer.Viewer{
					Sender: sender,
				},
			},
		},
	}

	// ignore context cancellation error: it does not matter how it was cancelled
	appIO.LoopReceiver(ctx, receiver, func(message []byte) {
		controller.HandleMessage(ctx, message)
	})
}

func setupRecvMessageUsecase(
	ctx context.Context,
	receiver appIO.Receiver,
	sender appIO.Sender,
) {
	controller := recvMsgController.Controller{
		Viewer: recvMsgViewer.Viewer{
			Sender: sender,
		},
	}

	// ignore context cancellation error: it does not matter how it was cancelled
	_ = appIO.LoopReceiver(ctx, receiver, func(message []byte) {
		controller.HandleMessage(ctx, message)
	})
}
