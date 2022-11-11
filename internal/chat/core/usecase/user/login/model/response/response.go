package response

type Model struct {
	UserID    UserID
	SessionID SessionID
}

type (
	UserID    string
	SessionID string
)
