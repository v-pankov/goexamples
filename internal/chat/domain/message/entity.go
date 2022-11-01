package message

import (
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		ID     ID
		UserID user.ID
		RoomID room.ID

		Text string
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

func (m Entity) Validate() error {
	if err := m.ID.Validate(); err != nil {
		return fmt.Errorf("id: %w", err)
	}

	if err := m.UserID.Validate(); err != nil {
		return fmt.Errorf("user id: %w", err)
	}

	if err := m.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	return nil
}
