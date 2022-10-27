package entity

import (
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Message struct {
		entity.Entity

		ID     MessageID
		UserID UserID
		RoomID RoomID

		Text string
	}

	MessageID string
)

var (
	ErrEmptyMessageID = errors.New("message id is empty")
)

func (id MessageID) Validate() error {
	if id == "" {
		return ErrEmptyMessageID
	}
	return nil
}

func (m Message) Validate() error {
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
