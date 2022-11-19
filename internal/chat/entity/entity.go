package entity

import (
	"time"
)

type (
	Message struct {
		ID       MessageID
		Contents MessageContents

		CreatedAt time.Time
	}

	MessageID       int64
	MessageContents []byte
)
