// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

type HubClient interface {
	Line() chan []byte
}

type Hub struct {
	// Registered clients.
	clients map[HubClient]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan HubClient

	// Unregister requests from clients.
	unregister chan HubClient
}

func (h *Hub) Register(c HubClient) {
	h.register <- c
}

func (h *Hub) Unregister(c HubClient) {
	h.unregister <- c
}

func (h *Hub) Broadcast(msg []byte) {
	h.broadcast <- msg
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan HubClient),
		unregister: make(chan HubClient),
		clients:    make(map[HubClient]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Line())
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Line() <- message:
				default:
					close(client.Line())
					delete(h.clients, client)
				}
			}
		}
	}
}
