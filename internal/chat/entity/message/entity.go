package message

import (
	"time"
)

type (
	Entity struct {
		ID   ID
		Text string

		CreatedAt time.Time
	}

	ID int64
)
