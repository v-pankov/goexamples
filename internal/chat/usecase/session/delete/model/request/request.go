package request

// Context is a part of use case request model port but consideted valid
type Context struct {
}

// Model is is a use case request model
type Model struct {
	SessionID SessionID
}

type SessionID string
