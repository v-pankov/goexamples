package event

import (
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity/event"
	"github.com/vdrpkv/goexamples/internal/chat/entity/event/user/data"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
)

type (
	Entity struct {
		EventID   event.ID
		SessionID session.ID

		Type Type
		Data Data

		CreatedAt time.Time
		DeletedAt *time.Time
	}

	Type string
	Data struct {
		New    *data.New
		Delete *data.Delete
	}
)

const (
	New    Type = "new"
	Delete Type = "delete"
)
