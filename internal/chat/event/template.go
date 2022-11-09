package event

import "time"

type Template[Header any, Type any, Data any] struct {
	Header Header
	Type   Type
	Data   Data
	Time   time.Time
}
