package message

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/event"
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		EventID   event.ID
		MessageID message.ID

		Type Type
	}

	Type string
)

const (
	Created Type = "created"
	Deleted Type = "deleted"
	Updated Type = "updated"
)
