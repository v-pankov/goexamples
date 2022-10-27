package entity

import (
	"errors"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Session struct {
		entity.Entity

		ID     SessionID
		UserID UserID
	}

	SessionID string
)

var (
	ErrEmptySessionID = errors.New("session id is empty")
)

func (id SessionID) Validate() error {
	if id == "" {
		return ErrEmptySessionID
	}
	return nil
}
