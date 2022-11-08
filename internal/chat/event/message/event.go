package event

import (
	"github.com/vdrpkv/goexamples/internal/pkg/event"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type (
	Event struct {
		event.Event

		ID   ID
		Type Type
		Data Data
	}

	ID   int64
	Type string
	Data struct {
		UserID      user.ID
		SessionID   session.ID
		MessageID   message.ID
		MessageText string
	}
)

const (
	New    Type = "new"
	Edit   Type = "edit"
	Delete Type = "delete"
)
