package response

import "time"

// Model is a use case response model
type Model struct {
	Messages <-chan *Message
}

type Message struct {
	SessionID SessionID
	Text      string
	CreatedAt time.Time
}

type SessionID string
