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
