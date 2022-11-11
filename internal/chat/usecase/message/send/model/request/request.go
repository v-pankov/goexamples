package request

import "time"

// Context is a part of use case request model port but consideted valid
type Context struct {
	UserID    UserID
	SessionID SessionID
}

// Model is is a use case request model
type Model struct {
	UserID      UserID
	SessionID   SessionID
	MessageID   MessageID
	MessageText MessageText
	CreatedAt   time.Time
}

type (
	UserID      string
	SessionID   string
	MessageID   int64
	MessageText string
)
