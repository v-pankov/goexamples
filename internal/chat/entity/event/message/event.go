package event

import (
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity/event"
	"github.com/vdrpkv/goexamples/internal/chat/entity/event/message/data"
	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
)

type (
	Entity struct {
		EventID   event.ID
		MessageID message.ID

		Type Type
		Data Data

		CreatedAt time.Time
		DeletedAt *time.Time
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
