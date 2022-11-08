package event

import (
	"github.com/vdrpkv/goexamples/internal/chat/entity/event"
	"github.com/vdrpkv/goexamples/internal/chat/entity/event/message/data"
	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
)

type (
	Event struct {
		EventID   event.ID
		MessageID message.ID

		Type Type
		Data Data
	}

	Type string
	Data struct {
		New    *data.New
		Edit   *data.Edit
		Delete *data.Delete
	}
)

const (
	New    Type = "new"
	Edit   Type = "edit"
	Delete Type = "delete"
)
