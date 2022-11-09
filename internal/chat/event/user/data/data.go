package data

import "github.com/vdrpkv/goexamples/internal/chat/event"

type Created struct {
	UserName event.UserID
}

type Updated struct {
	UserName event.UserName
}

type Deleted struct {
}
