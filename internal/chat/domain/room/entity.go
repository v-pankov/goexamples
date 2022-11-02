package room

import (
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		ID     ID
		UserID user.ID

		Name string
	}

	ID   string
	Name string
)

func (id ID) String() string {
	return string(id)
}

func (n Name) String() string {
	return string(n)
}
