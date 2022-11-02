package user

import (
	"errors"
	"strings"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"
)

type (
	Entity struct {
		entity.Entity

		ID ID

		Name Name
	}

	ID   string
	Name string
)

var (
	ErrEmptyID   = errors.New("id is empty")
	ErrEmptyName = errors.New("name is empty")
)

func (id ID) Validate() error {
	if id == "" {
		return ErrEmptyID
	}
	return nil
}

func (n Name) Validate() error {
	if n == "" {
		return ErrEmptyName
	}

	if len(strings.TrimSpace(string(n))) == 0 {
		return ErrEmptyName
	}

	return nil
}
