package response

import "time"

type Model struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       time.Time
}
