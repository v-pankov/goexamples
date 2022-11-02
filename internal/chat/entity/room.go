package entity

import (
	"errors"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Room struct {
		entity.Entity

		ID     RoomID
		UserID UserID

		Name string
	}

	RoomName string

	RoomID string
)

var (
	ErrEmptyRoomID   = errors.New("room id is empty")
	ErrEmptyRoomName = errors.New("room name is empty")
)

func (id RoomID) Validate() error {
	if id == "" {
		return ErrEmptyRoomID
	}
	return nil
}

func (n RoomName) Validate() error {
	if n == "" {
		return ErrEmptyRoomName
	}
	return nil
}
