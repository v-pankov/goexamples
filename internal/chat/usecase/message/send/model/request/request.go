package request

// Context is a part of use case request model port but consideted valid
type Context struct {
	UserID    UserID
	SessionID SessionID
}

// Model is is a use case request model
type Model struct {
	MessageText string
}

type (
	UserID    string
	SessionID string
)
