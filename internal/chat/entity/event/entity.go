package event

import "time"

type (
	Entity struct {
		ID   ID
		Type Type

		FiredAt   time.Time
		SavedAt   time.Time
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
