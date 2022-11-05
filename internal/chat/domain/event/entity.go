package event

import (
	"time"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		ID   ID
		Type Type

		FiredAt time.Time
	}

	ID   string
	Type string
)

const (
	Message Type = "message"
	User    Type = "user"
)
