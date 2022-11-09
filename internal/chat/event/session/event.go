package session

import (
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/event"
	"github.com/vdrpkv/goexamples/internal/chat/event/user/data"
)

type (
	Event struct {
		UserID    event.UserID
		SessionID event.SessionID

		Type Type
		Data Data

		Time time.Time
	}

	Type string
	Data struct {
		Created *data.Created
		Deleted *data.Deleted
	}
)

const (
	Created Type = "created"
	Deleted Type = "deleted"
)
