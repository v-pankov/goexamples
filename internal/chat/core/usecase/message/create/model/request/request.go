package request

// Context is a part of usecase request model consideted valid
type Context struct {
	UserID UserID
}

type Model struct {
	MessageText string
}

type (
	UserID string
)
