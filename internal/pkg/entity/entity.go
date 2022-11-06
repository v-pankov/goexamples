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

func (e *Entity) UpdateTime() {
	e.UpdatedAt = time.Now()
}

func New() Entity {
	timeNow := time.Now()
	return Entity{
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}
