package event

import (
	"time"

	"github.com/vdrpkv/goexamples/internal/chat/entity/event"
	"github.com/vdrpkv/goexamples/internal/chat/entity/event/user/data"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
)

type (
	Entity struct {
		EventID event.ID
		UserID  user.ID

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
