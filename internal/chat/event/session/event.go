package session

import (
	"github.com/vdrpkv/goexamples/internal/chat/event"
	"github.com/vdrpkv/goexamples/internal/chat/event/user/data"
)

type (
	Event event.Template[Header, Type, Data]

	Header struct {
		UserID    event.UserID
		SessionID event.SessionID
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
