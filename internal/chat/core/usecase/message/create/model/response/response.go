package response

import "time"

type Model struct {
	Message Message
}

type (
	Message struct {
		MessageID   MessageID
		MessageText MessageText
		CreatedAt   time.Time
	}

	MessageID   int64
	MessageText string
)
