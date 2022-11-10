package response

import "time"

// Model is a use case response model
type Model struct {
	Message Message
}

type (
	Message struct {
		UserID      UserID
		SessionID   SessionID
		MessageID   MessageID
		MessageText MessageText
		CreatedAt   time.Time
	}

	UserID      string
	SessionID   string
	MessageID   int64
	MessageText string
)
