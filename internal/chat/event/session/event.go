package session

import (
	"github.com/vdrpkv/goexamples/internal/pkg/event"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/event/user/data"
)

type (
	Event struct {
		event.Event

		UserID    user.ID
		SessionID session.ID

		Type Type
		Data Data
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
