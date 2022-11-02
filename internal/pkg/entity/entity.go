package entity

import "time"

type Entity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (e *Entity) Deleted() bool {
	return e.DeletedAt == nil
}
