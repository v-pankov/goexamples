package session

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		ID     ID
		UserID user.ID
	}

	ID string
)

func (id ID) String() string {
	return string(id)
}
