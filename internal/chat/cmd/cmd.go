package cmd

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/vdrpkv/goexamples/internal/chat/core"
	"github.com/vdrpkv/goexamples/internal/chat/infrastructure/pubsub"
	"github.com/vdrpkv/goexamples/internal/chat/infrastructure/websocket"

	pubsubInmem "github.com/vdrpkv/goexamples/internal/chat/infrastructure/pubsub/inmem"

	inmemRepo "github.com/vdrpkv/goexamples/internal/chat/infrastructure/repository/inmem"
	usecaseMessageSendInmemRepo "github.com/vdrpkv/goexamples/internal/chat/infrastructure/repository/inmem/message/send"

	usecaseMessageRecv "github.com/vdrpkv/goexamples/internal/chat/usecase/message/recv"
	usecaseMessageSend "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send"
)

var addr = flag.String("addr", ":8080", "http service address")

func Run() {
	flag.Parse()

	ctx := context.Background()

	pub := pubsubInmem.New()
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
	receiver core.Receiver,
	sender core.Sender,
	repo *inmemRepo.InMem,
) {
	var (
		usecaseRepo = usecaseMessageSendInmemRepo.New(repo)
		processor   = usecaseMessageSend.NewProcessor(usecaseRepo)
		viewer      = usecaseMessageSend.NewViewer(
			sender,
			core.ErrorHandlerFunc(
				func(ctx context.Context, err error) {
					log.Println(err)
				},
			),
		)
		presenter  = usecaseMessageSend.NewPresenter(viewer)
		interactor = usecaseMessageSend.NewInteractor(
			processor,
			presenter,
			core.ErrorHandlerFunc(func(ctx context.Context, err error) {
				log.Println(err)
			}),
		)
		controller = usecaseMessageSend.NewController(receiver, interactor)
	)
	controller.Run(ctx)
}

func setupRecvMessageUsecase(
	ctx context.Context,
	receiver core.Receiver,
	sender core.Sender,
) {
	var (
		viewer = usecaseMessageRecv.NewViewer(
			sender,
			core.ErrorHandlerFunc(
				func(ctx context.Context, err error) {
					log.Println(err)
				},
			),
		)
		controller = usecaseMessageRecv.NewController(
			receiver,
			viewer,
			core.ErrorHandlerFunc(
				func(ctx context.Context, err error) {
					log.Println(err)
				},
			),
		)
	)
	controller.Run(ctx)
}
