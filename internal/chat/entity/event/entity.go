package event

import "time"

type (
	Entity struct {
		ID   ID
		Type Type

		FiredAt   time.Time
		CreatedAt time.Time
		DeletedAt *time.Time
	}

	ID   string
	Type string
)

const (
	Message Type = "message"
	Session Type = "session"
	User    Type = "user"
)
