package entity

import (
	"errors"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	User struct {
		entity.Entity

		ID UserID

		Name UserName
	}

	UserName string

	UserID string
)

var (
	ErrEmptyUserID   = errors.New("user id is empty")
	ErrEmptyUserName = errors.New("user name is empty")
)

func (id UserID) Validate() error {
	if id == "" {
		return ErrEmptyUserID
	}
	return nil
}

func (n UserName) Validate() error {
	if n == "" {
		return ErrEmptyUserName
	}
	return nil
}
