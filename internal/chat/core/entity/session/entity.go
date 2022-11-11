package session

import (
	"github.com/vdrpkv/goexamples/internal/chat/core/entity/user"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		ID     ID
		UserID user.ID

		Active bool
	}

	ID string
)

func (id ID) String() string {
	return string(id)
}
