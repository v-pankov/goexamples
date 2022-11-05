package user

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/event"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		EventID   event.ID
		SessionID session.ID

		Type Type
	}

	Type string
)

const (
	Login  Type = "login"
	Logout Type = "logout"
)
