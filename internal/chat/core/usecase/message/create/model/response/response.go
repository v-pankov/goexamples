package response

import "time"

type Model struct {
	Message Message
}

type (
	Message struct {
		UserID      UserID
		MessageID   MessageID
		MessageText MessageText
		CreatedAt   time.Time
	}

	UserID      string
	MessageID   int64
	MessageText string
)
