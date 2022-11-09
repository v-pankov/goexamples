package response

// Model is a use case response model
type Model struct {
	SessionID SessionID
}

// SessionID is used to differ one user session from another
type SessionID string
