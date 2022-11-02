package user

import (
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

func (i ID) String() string {
	return string(i)
}

func (n Name) String() string {
	return string(n)
}
