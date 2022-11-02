package session

import (
	"errors"

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

var (
	ErrEmptyID = errors.New("id is empty")
)

func (id ID) Validate() error {
	if id == "" {
		return ErrEmptyID
	}
	return nil
}
