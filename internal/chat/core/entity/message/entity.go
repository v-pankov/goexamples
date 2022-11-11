package message

import (
	"github.com/vdrpkv/goexamples/internal/pkg/entity"

	"github.com/vdrpkv/goexamples/internal/chat/core/entity/user"
)

type (
	Entity struct {
		entity.Entity

		ID     ID
		UserID user.ID
		Text   string
	}

	ID int64
)

func (id ID) Int64() int64 {
	return int64(id)
}
