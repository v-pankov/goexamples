package message

import (
	"github.com/vdrpkv/goexamples/internal/chat/event"
	"github.com/vdrpkv/goexamples/internal/chat/event/message/data"
)

type (
	Event event.Template[Header, Type, Data]

	Header struct {
		UserID    event.UserID
		SessionID event.SessionID
		MessageID event.MessageID
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
