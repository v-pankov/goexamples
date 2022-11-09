package response

// Model is a use case response model
type Model struct {
	UserID UserID
}

// UserID is used to differ one user from another
type UserID string
