package request

// Context is a part of use case request model port but consideted valid
type Context struct {
	SessionID SessionID
}

// Model is is a use case request model
type Model struct {
	MessageText string
}

// SessionID is used to differ one user session from another
type SessionID string
