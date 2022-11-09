package message

import (
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/event"
	"github.com/vdrpkv/goexamples/internal/chat/event/user/data"
)

type (
	Event struct {
		UserID    event.UserID
		SessionID event.SessionID
		MessageID event.MessageID

		Type Type
		Data Data

		Time time.Time
	}

	Type string
	Data struct {
		Created *data.Created
		Updated *data.Updated
		Deleted *data.Deleted
	}
)

const (
	Created Type = "created"
	Updated Type = "updated"
	Deleted Type = "deleted"
)
