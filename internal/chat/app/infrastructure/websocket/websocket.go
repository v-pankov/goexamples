package websocket

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	appIO "github.com/vdrpkv/goexamples/internal/chat/app/io"
)

const (
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
)

var (
	newline = []byte{'\n'}
)

type (
	MessageType string
)

const (
	Binary MessageType = "binary"
	Text   MessageType = "text"
)

var (
	ErrUnexpectedMessageType = errors.New("unexpected message type")
)

// readPump pumps messages from the websocket connection.
func readPump(ctx context.Context, wsMessageType int, conn *websocket.Conn, sink chan<- []byte) <-chan error {
	const (
		// Maximum message size allowed from peer.
		maxMessageSize = 512
	)

	var (
		space = []byte{' '}
	)

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	var (
		errChan = make(chan error)
	)

	go func() {
		defer close(errChan)
		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				messageType, messageBytes, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(
						err,
						websocket.CloseGoingAway,
						websocket.CloseAbnormalClosure,
					) {
						errChan <- fmt.Errorf("unexpected close: %w", err)
					} else {
						errChan <- fmt.Errorf("read message: %w", err)
					}
					return
				}

				if messageType != wsMessageType {
					errChan <- ErrUnexpectedMessageType
					return
				}

				messageBytes = bytes.TrimSpace(bytes.Replace(
					messageBytes, newline, space, -1,
				))

				sink <- messageBytes
			}
		}
	}()

	return errChan
}

// writePump pumps messages to the websocket connection.
func writePump(ctx context.Context, wsMessageType int, conn *websocket.Conn, source <-chan []byte) <-chan error {
	const (
		// Time allowed to write a message to the peer.
		writeWait = 10 * time.Second
	)

	var (
		errSourceClosed = errors.New("source is closed")
	)

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()

	var (
		errChan = make(chan error)
	)

	go func() {
		defer close(errChan)
		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			case messageBytes, ok := <-source:
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				if !ok {
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					errChan <- errSourceClosed
				}

				if err := conn.WriteMessage(wsMessageType, messageBytes); err != nil {
					errChan <- fmt.Errorf("write message: %w", err)
					return
				}
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					errChan <- fmt.Errorf("write ping message: %w", err)
					return
				}
			}
		}
	}()

	return errChan
}

type Handler struct {
	send chan []byte
	recv chan []byte
	err  chan error
}

func New() *Handler {
	return &Handler{
		send: make(chan []byte),
		recv: make(chan []byte),
	}
}

func (h *Handler) Serve(parentCtx context.Context, w http.ResponseWriter, r *http.Request, messageType MessageType) error {
	var wsMessageType int
	switch messageType {
	case Binary:
		wsMessageType = websocket.BinaryMessage
	case Text:
		wsMessageType = websocket.TextMessage
	default:
		return ErrUnexpectedMessageType
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return fmt.Errorf("upgrade: %w", err)
	}

	ctx, ctxCancel := context.WithCancel(parentCtx)

	var (
		readErr  = readPump(ctx, wsMessageType, conn, h.recv)
		writeErr = writePump(ctx, wsMessageType, conn, h.send)
	)

	go func() {
		defer func() {
			ctxCancel()
			close(h.err)
		}()

		select {
		case err := <-readErr:
			h.err <- err
		case err := <-writeErr:
			h.err <- err
		}
	}()

	return nil
}

func (h *Handler) Receive(context.Context) <-chan []byte {
	return h.recv
}

func (h *Handler) Send(_ context.Context, bytes []byte) error {
	h.send <- bytes
	return nil
}

var (
	_ appIO.Receiver = (*Handler)(nil)
	_ appIO.Sender   = (*Handler)(nil)
)
