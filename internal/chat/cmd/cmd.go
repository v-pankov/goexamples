package cmd

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"

	inmemPubSub "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/pubsub/inmem"
	inmemRepo "github.com/vdrpkv/goexamples/internal/chat/app/infrastructure/repository/inmem"

	serverHome "github.com/vdrpkv/goexamples/internal/chat/cmd/serve/home"
	serverWebsocket "github.com/vdrpkv/goexamples/internal/chat/cmd/serve/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

func Run() {
	flag.Parse()

	var (
		ctx       = context.Background()
		pub       = inmemPubSub.New()
		inmemRepo = inmemRepo.New()
		wg        sync.WaitGroup
	)

	go pub.Run(ctx)

	http.HandleFunc("/", serverHome.Handler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWebsocket.Handler(ctx, w, r, &wg, pub, inmemRepo)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	wg.Wait()
}
